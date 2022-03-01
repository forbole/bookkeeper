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
	vsCurrency:="USD"

	

	// getting balance for the address
	// May need to query from graphql if we want to get exact balance
	// Like what X does?
	// For now just using the price
	balances,err:=ParseBalanceSheet(coin,vsCurrency,CG)
	if err!=nil{
		panic(err)
	}

	
	// Output the .csv file contains the
	// schema 
	// date, coin price, account balance
	outputcsv:=balances.GetCSV()
	fmt.Println(outputcsv)
}

func ParseBalanceSheet(c types.Coin,vsCurrency string,CG *coingecko.Client)(types.Balances,error){
	var balances types.Balances
	
		// getting coinprice fact
		days:=time.Since(c.StartDate).Hours()/24
		// DAILTY data will be used for duration above 90 days.
		coindata,err:=CG.CoinsIDMarketChart(c.CoinType,vsCurrency,strconv.FormatFloat(days, 'f', -1, 64))
		if err!=nil{
			return nil,err
		}
		
		// convert it to []types.Balance
		countmonth:=1
		countday:=0
		prices:=*coindata.Prices
		for countday<len(prices){
			


			balances = append(balances,types.NewBalance(&c,prices[countday][1],
				c.StartDate.Add(time.Hour*time.Duration(24*countday)), 
				prices[countday][1]*float32(c.Quantity),vsCurrency))
			
			countmonth+=1

			if countmonth==13{
				countmonth=1
			}
			if countmonth%2==0{
				countday+=30
				}else{
					countday+=31
				}
		}

		return balances,nil
}