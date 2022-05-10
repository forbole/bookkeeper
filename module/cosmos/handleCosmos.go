package cosmos

import (
	"fmt"
	"io/ioutil"

	tables "github.com/forbole/bookkeeper/module/cosmos/tables"
	types "github.com/forbole/bookkeeper/types"
)

// HandleCosmos process all the chain in the struct.
// Make a .csv file at "." and return the relative path
func HandleCosmosMonthyReport(individualChains []types.IndividualChain, vsCurrency string, outputFolder string,
	period types.Period) ([]string, error) {
	var filenames []string

	for _, data := range individualChains {
		entries, err := tables.GetMonthyReport(data,
			period)
		if err != nil {
			return nil, err
		}

		// Writ .csv to "."
		for _, e := range entries {
			outputcsv, err := e.Rows.GetCSVConvertedPrice(data.Denom, vsCurrency)
			if err != nil {
				return nil, err
			}
			outputcsv2 := e.Rows.GetCSV(0)
			outputcsv3, err := e.Rows.GetMonthyCSVConvertedPrice(data.Denom, vsCurrency)
			if err != nil {
				return nil, err
			}

			if err != nil {
				return nil, err
			}
			//fmt.Println(outputcsv)
			filename := fmt.Sprintf("%s/%s_%s_monthy_report_value.csv", outputFolder, data.ChainName, e.Address)
			err = ioutil.WriteFile(filename, []byte(outputcsv), 0600)
			if err != nil {
				return nil, err
			}
			filenames = append(filenames, filename)

			filename2 := fmt.Sprintf("%s/%s_%s_monthy_report.csv", outputFolder, data.ChainName, e.Address)
			err = ioutil.WriteFile(filename2, []byte(outputcsv2), 0600)
			if err != nil {
				return nil, err
			}
			filenames = append(filenames, filename2)

			filename3 := fmt.Sprintf("%s/%s_%s_monthy_report_convert_price_only.csv", outputFolder, data.ChainName, e.Address)
			err = ioutil.WriteFile(filename3, []byte(outputcsv3), 0600)
			if err != nil {
				return nil, err
			}
			filenames = append(filenames, filename3)
		}

	}
	return filenames, nil
}

func HandleTxsTable(individualChains []types.IndividualChain, outputFolder string,
	period types.Period) ([]string, error) {
	var filenames []string
	for _, detail := range individualChains {
		entries, err := tables.GetTxs(detail, period.From)
		if err != nil {
			return nil, err
		}
		for _, e := range entries {
			outputcsv := e.Rows.GetCSV()
			if err != nil {
				return nil, err
			}
			//fmt.Println(outputcsv)
			filename := fmt.Sprintf("%s/%s_txs.csv", outputFolder, e.Address)
			err = ioutil.WriteFile(filename, []byte(outputcsv), 0600)
			if err != nil {
				return nil, err
			}
			filenames = append(filenames, filename)
		}
	}
	return filenames, nil
}

func HandleRewardCommissionTable(individualChains []types.IndividualChain,
	outputFolder string, period types.Period) ([]string, error) {
	var filenames []string
	for _, detail := range individualChains {
		txs, err := tables.GetTxs(detail, period.From)
		if err != nil {
			return nil, err
		}

		for _, tx := range txs {
			e, err := tables.GetRewardCommission(tx)
			if err != nil {
				return nil, err
			}

			outputcsv := e.Rows.GetCSV()
			if err != nil {
				return nil, err
			}
			//fmt.Println(outputcsv)
			filename := fmt.Sprintf("%s/%s_reward_commission.csv", outputFolder, e.Address)
			err = ioutil.WriteFile(filename, []byte(outputcsv), 0600)
			if err != nil {
				return nil, err
			}

			filenames = append(filenames, filename)
		}

	}

	return filenames, nil
}
