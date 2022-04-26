package cosmos

import (
	"fmt"
	"io/ioutil"
	"time"

	tables "github.com/forbole/bookkeeper/module/cosmos/tables"
	types "github.com/forbole/bookkeeper/types"


)

// HandleCosmos process all the chain in the struct.
// Make a .csv file at "." and return the relative path
func HandleCosmosMonthyReport(individualChains []types.IndividualChain,vsCurrency string,outputFolder string)([]string,error){
	var filenames []string

	for _,data :=range individualChains{
		entries, err := tables.GetMonthyReport(data,
			time.Date(2022,time.January,1,1,0,0,0,time.UTC))
		if err != nil {
			return nil,err
		}

		// Writ .csv to "." 
		for _ ,e:=range entries{
			//outputcsv,err := e.Rows.GetCSVConvertedPrice(data.Denom,vsCurrency)
			outputcsv := e.Rows.GetCSV(0)
			if err!=nil{
				return nil,err
			}
			fmt.Println(outputcsv)
			filename := fmt.Sprintf("%s/%s_monthy_report.csv", outputFolder,e.Address)
			err = ioutil.WriteFile(filename, []byte(outputcsv), 0777)
			if err != nil {
				return nil,err
			}
			filenames = append(filenames, filename)
		}

	}
	return filenames,nil
}

func HandleTxsTable(individualChains []types.IndividualChain,outputFolder string)([]string,error){
	var filenames []string
	for _,detail:=range individualChains{
		entries,err:=tables.GetTxs(detail)
		if err!=nil{
			return nil,err
		}
		for _ ,e:=range entries{
			outputcsv:= e.Rows.GetCSV()
			if err!=nil{
				return nil,err
			}
			fmt.Println(outputcsv)
			filename := fmt.Sprintf("%s/%s_txs.csv", outputFolder,e.Address)
			err = ioutil.WriteFile(filename, []byte(outputcsv), 0777)
			if err != nil {
				return nil,err
			}
			filenames = append(filenames, filename)
		}
	}
	return filenames,nil
}

func HandleRewardCommissionTable(individualChains []types.IndividualChain,outputFolder string)([]string,error){
	var filenames []string
	for _,detail:=range individualChains{
		txs,err:=tables.GetTxs(detail)
		if err!=nil{
			return nil,err
		}

		for _,tx:=range txs{
			e,err:=tables.GetRewardCommission(tx)
			if err!=nil{
				return nil,err
			}

				outputcsv:= e.Rows.GetCSV()
				if err!=nil{
					return nil,err
				}
				fmt.Println(outputcsv)
				filename := fmt.Sprintf("%s/%s_reward_commission.csv", outputFolder,e.Address)
				err = ioutil.WriteFile(filename, []byte(outputcsv), 0777)
				if err != nil {
					return nil,err
				}

				filenames = append(filenames, filename)
			}

		}
		
	return filenames,nil
}