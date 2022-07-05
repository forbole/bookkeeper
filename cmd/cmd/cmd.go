package parse

import (
	"fmt"
	"os"
	"path/filepath"

	//"time"

	"io/fs"

	//"net/http"

	//"github.com/forbole/bookkeeper/balancesheet"
	//"github.com/cosmos/cosmos-sdk/client"

	"github.com/forbole/bookkeeper/email"
	"github.com/forbole/bookkeeper/module/cosmos"
	"github.com/forbole/bookkeeper/module/elrond"
	"github.com/forbole/bookkeeper/module/flow"
	"github.com/forbole/bookkeeper/module/substrate"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/forbole/bookkeeper/utils"

	"github.com/joho/godotenv"

	//"google.golang.org/grpc"

	//coingecko "github.com/superoo7/go-gecko/v3"

	"github.com/spf13/cobra"
)

const (
	flagInputJsonPath = "input_json_path"
	flagOutputFolder  = "output_folder"
	flagLogLevel = "log_level"
)

// ParseCmd returns the command that should be run when we want to start parsing a chain state.
func ParseCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "bookkeeper",
		Short: "Start parsing the blockchain data",
		RunE:  Execute,
	}
	cmd.Flags().String(flagInputJsonPath, "./input.json", "The path that the input file should read from")
	cmd.Flags().String(flagOutputFolder, "./output", "The path output .csv file sit at")
	cmd.Flags().Int8(flagLogLevel, int8(-1), "log level that output")

	return &cmd
}

func Execute(cmd *cobra.Command, arg []string) error {
	jsonPath, _ := cmd.Flags().GetString(flagInputJsonPath)
	outputFile, _ := cmd.Flags().GetString(flagOutputFolder)
	logLevel, _ := cmd.Flags().GetInt8(flagLogLevel)

    zerolog.SetGlobalLevel(zerolog.Level(logLevel))

	err := godotenv.Load()
	if err != nil {
		return err
	}

	// This returns an *os.FileInfo type
	fileInfo, err := os.Stat(jsonPath)
	if err != nil {
		return err
	}
fmt.Println(fileInfo.IsDir())
	// IsDir is short for fileInfo.Mode().IsDir()
	if fileInfo.IsDir() {
	// file is a directory
	err =filepath.Walk(jsonPath,func(path string, _ fs.FileInfo, _ error)error{
			fmt.Println(path)
			if path==jsonPath{
				return nil
			}		
			err=handleSingleFile(path,outputFile)
				if err!=nil{
					log.Error().Msgf("cannot parse:%s file:%s",err,path)
					return nil
				}
				return nil
		})
	if err!=nil{
		return err
	}
	} else {
		err=handleSingleFile(jsonPath,outputFile)
		if err!=nil{
			return err
		}
	}

	return nil

}

// handleSingleFile handle a json file
func handleSingleFile(jsonPath string,outputFile string)error{
	data, err := utils.ImportJsonInput(jsonPath)
	if err != nil {
		return err
	}

	// make output directory
	if _, err := os.Stat(outputFile); os.IsNotExist(err) {
		if err := os.MkdirAll(outputFile, os.ModePerm); err != nil {
			return err
		}
	}

	//fmt.Println(*data)

	//inputfile:=[]string{"bitcoin.csv","ethereum.csv"}

	var filenames []string
	for _ , chain:=range data.Chains {
		files, err := cosmos.HandleRewardPriceTable(chain, data.VsCurrency, outputFile, data.Period)
		if err != nil {
			return err
		}

		filenames = append(filenames, files...)

	}

	if data.Flow.Db.Port != 0 {
		flowfile, err := flow.HandleRewardTable(data.Flow, data.VsCurrency, data.Period)
		if err != nil {
			return err
		}
		filenames = append(filenames, flowfile...)

	}

	for _, sub := range data.Substrate {
		substratefile, err := substrate.Handle(sub, data.VsCurrency, outputFile, data.Period)
		if err != nil {
			return err
		}

		filenames = append(filenames, substratefile...)
	}

	if data.Elrond.Addresses != nil {
		file, err := elrond.HandleTx(data.Elrond, data.Period, outputFile, data.VsCurrency)
		if err != nil {
			return err
		}
		filenames = append(filenames, file...)

	}

	err = email.SendEmail(data.EmailDetails, filenames)
	if err != nil {
		return err
	}
	return nil
}
