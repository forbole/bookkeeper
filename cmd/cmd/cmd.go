package parse

import (
	"fmt"
	"os"

	//"time"

	"io/ioutil"

	//"net/http"

	//"github.com/forbole/bookkeeper/balancesheet"
	//"github.com/cosmos/cosmos-sdk/client"

	"github.com/forbole/bookkeeper/email"
	"github.com/forbole/bookkeeper/module/cosmos"
	"github.com/forbole/bookkeeper/module/flow"
	"github.com/forbole/bookkeeper/module/subtrate"

	"github.com/forbole/bookkeeper/utils"

	"github.com/joho/godotenv"

	"github.com/forbole/bookkeeper/types"

	//"google.golang.org/grpc"

	//coingecko "github.com/superoo7/go-gecko/v3"

	"github.com/spf13/cobra"
)

const (
	flagInputJsonPath = "input_json_path"
	flagOutputFolder  = "output_folder"
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
	return &cmd
}

func Execute(cmd *cobra.Command, arg []string) error {
	jsonPath, _ := cmd.Flags().GetString(flagInputJsonPath)
	outputFile, _ := cmd.Flags().GetString(flagOutputFolder)

	err := godotenv.Load()
	if err != nil {
		return err
	}

	data, err := utils.ImportJsonInput(jsonPath)
	if err != nil {
		return err
	}
	//fmt.Println(*data)

	//inputfile:=[]string{"bitcoin.csv","ethereum.csv"}

	// make output directory
	if _, err := os.Stat(outputFile); os.IsNotExist(err) {
		if err:=os.MkdirAll(outputFile,os.ModePerm);err!=nil{
			return err
		}
	}

	var filenames []string

	files2, err := cosmos.HandleRewardPriceTable(data.Chains, data.VsCurrency, outputFile, data.Period)
	if err != nil {
		return err
	}
	filenames = append(filenames, files2...)
	/*
			files3,err:=cosmos.HandleTxsTable(chain.Details,outputFile,data.Period)
			if err!=nil{
				return err
			}
			filenames = append(filenames, files3...)

		files4,err:=cosmos.HandleRewardCommissionTable(chain.Details,outputFile,data.Period)
			if err!=nil{
				return err
			}
			filenames = append(filenames, files4...) */

	for _, chain := range data.Subtrate {
		file3, err := subtrate.Handle(chain,data.VsCurrency,outputFile,data.Period)
		if err != nil {
			return err
		}
		filenames = append(filenames, file3...)

	}

	flowfile, err := flow.HandleRewardTable(data.Flow, data.VsCurrency, data.Period)
	if err != nil {
		return err
	}

	filenames = append(filenames, flowfile...)

	err = email.SendEmail(data.EmailDetails, filenames)
	if err != nil {
		return err
	}

	/* grpcConn, err := grpc.Dial(data.Chains[0].Details[0].GrpcEndpoint, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer grpcConn.Close() */

	//coingecko
	/* httpClient := &http.Client{
			Timeout: time.Second * 10,
		}
		cg := coingecko.NewClient(httpClient)

		// Import coin info
		// Assume the dateQuantity pair is always by time asc
		coin := types.NewCoin("bitcoin", []types.DateQuantity{
			types.NewDateQuantity(5.1234,
				time.Date(2020, time.January, 28, 0, 0, 0, 0, time.UTC)),
			types.NewDateQuantity(15.5642,
				time.Date(2020, time.August, 28, 0, 0, 0, 0, time.UTC)),
		})

		eth := types.NewCoin("ethereum", []types.DateQuantity{
			types.NewDateQuantity(30.4501,
				time.Date(2020, time.January, 28, 0, 0, 0, 0, time.UTC)),
			types.NewDateQuantity(16.4564,
				time.Date(2020, time.December, 28, 0, 0, 0, 0, time.UTC)),
		})

		vsCurrency := "USD"

		// getting balance for the address
		// May need to query from graphql if we want to get exact balance
		// Like what X does?
		// For now just using the price
		balances, err := balancesheet.ParseBalanceSheet(coin, vsCurrency, cg)
		if err != nil {
			return err
		}

		totalBalance,err :=balancesheet.TotalValueBalanceSheet([]types.Coin{
			coin,eth,
		},vsCurrency,cg)

		ethBalance,err := balancesheet.ParseBalanceSheet(eth, vsCurrency, cg)
		if err!=nil{
			return err
		}

		// Output the .csv file contains the
		// schema
		// date, coin price, account balance
		outputcsv := balances.GetCSV()
		//fmt.Println(outputcsv)
		err = ioutil.WriteFile(fmt.Sprintf("%s.csv",balances[0].Coin), []byte(outputcsv), 0600)
	    if err != nil {
	        return err
	    }

		totalCsv:=totalBalance.GetCSV()
		//fmt.Println(totalCsv)
		err= ioutil.WriteFile("totalValue.csv", []byte(totalCsv), 0600)
		if err!=nil{
			return err
		}

		err=OutputCsv(ethBalance)
		if err!=nil{
			return err
		} */
	return nil
}

func OutputCsv(b types.Balances) error {
	totalCsv := b.GetCSV()
	//fmt.Println(totalCsv)
	return ioutil.WriteFile(fmt.Sprintf("%s.csv", b[0].Coin), []byte(totalCsv), 0600)
}