package main

import (
	"fmt"
	"io/ioutil"
	"time"

	"net/http"

	"github.com/HarleyAppleChoi/bookeeper/types"
	"github.com/HarleyAppleChoi/bookeeper/balancesheet"

	coingecko "github.com/superoo7/go-gecko/v3"
)

func main() {
	//coingecko
	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}
	CG := coingecko.NewClient(httpClient)

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
	balances, err := balancesheet.ParseBalanceSheet(coin, vsCurrency, CG)
	if err != nil {
		panic(err)
	}
	
	totalBalance,err :=balancesheet.TotalValueBalanceSheet([]types.Coin{
		coin,eth,
	},vsCurrency,CG)

	// Output the .csv file contains the
	// schema
	// date, coin price, account balance
	outputcsv := balances.GetCSV()
	fmt.Println(outputcsv)
	err = ioutil.WriteFile(fmt.Sprintf("%s.csv",balances[0].Coin), []byte(outputcsv), 0777)
    if err != nil {
        panic(err)
    }

	err=OutputCsv(totalBalance)
	if err!=nil{
		panic(err)
	}
}

func OutputCsv(b types.Balances)error{
	totalCsv:=b.GetCSV()
	fmt.Println(totalCsv)
	return ioutil.WriteFile(fmt.Sprintf("%s.csv",b[0].Coin), []byte(totalCsv), 0777)
}