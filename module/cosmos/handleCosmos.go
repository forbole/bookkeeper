package cosmos

import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/forbole/bookkeeper/types"
)

// HandleCosmos process all the chain in the struct.
// Make a .csv file at "." and return the relative path
func HandleCosmosMonthyReport(individualChains []types.IndividualChain)([]string,error){
	var filenames []string

	for _,data :=range individualChains{
		entries, err := GetMonthyReport(data,
			time.Date(2022,time.January,1,1,0,0,0,time.UTC))
		if err != nil {
			return nil,err
		}

		// Writ .csv to "." 
		for _ ,e:=range entries{
			outputcsv,err := e.Rows.GetCSVConvertedPrice(6,data.ChainName,"USD")
			if err!=nil{
				return nil,err
			}
			fmt.Println(outputcsv)
			filename := fmt.Sprintf("%s.csv", e.Address)
			err = ioutil.WriteFile(filename, []byte(outputcsv), 0777)
			if err != nil {
				return nil,err
			}
			filenames = append(filenames, filename)
		}
	}

	return filenames,nil
}