package main

import (
	"fmt"
	"strconv"
	"time"

	"net/http"

	"github.com/HarleyAppleChoi/bookeeper/types"
	coingecko "github.com/superoo7/go-gecko/v3"
)

func main(){
	//coingecko
	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}
	CG := coingecko.NewClient(httpClient)

	// Import coin info
	coin:=types.NewCoin("bitcoin",5.1234,
	time.Date(2020,time.January,28,0,0,0,0,time.UTC))


	// getting coinprice fact
	days:=time.Since(coin.StartDate).Hours()/24
	// DAILTY data will be used for duration above 90 days.
	coindata,err:=CG.CoinsIDMarketChart(coin.CoinType,"USD",strconv.FormatFloat(days, 'f', -1, 64))
	if err!=nil{
		panic(err)
	}

	// getting balance for the address
	// May need to query from graphql if we want to get exact balance
	// Like what X does?
	// For now just using the price

	
	// Output the .csv file contains the
	// schema 
	// date, coin price, account balance
	countmonth:=1
	i:=0
	outputcsv:="date,coin, coin price, account balance\n"
	prices:=*coindata.Prices
	for i<len(prices){
		outputcsv+=fmt.Sprintf("%s,%s,%f,%f\n",
			coin.StartDate.Add(time.Hour*time.Duration(24*i)).String(), 
			coin.CoinType,
			prices[i][1],prices[i][1]*float32(coin.Quantity))
		
		countmonth+=1

		if countmonth==13{
			countmonth=1
		}
		if countmonth%2==0{
			i+=30
			}else{
				i+=31
			}
	}
	fmt.Println(outputcsv)
}