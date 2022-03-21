package balancesheet

import (
	"strconv"
	"time"

	"github.com/forbole/bookkeeper/types"
	coingecko "github.com/superoo7/go-gecko/v3"

)

// It parse single type of coin and return its balance sheet in certain period
func ParseBalanceSheet(c types.Coin, vsCurrency string, CG *coingecko.Client) (types.Balances, error) {
	var balances types.Balances

	// getting coinprice fact
	days := time.Since(c.DateQuantities[0].Date).Hours() / 24
	// DAILTY data will be used for duration above 90 days.
	coindata, err := CG.CoinsIDMarketChart(c.CoinType, vsCurrency, strconv.FormatFloat(days, 'f', -1, 64))
	if err != nil {
		return nil, err
	}

	// convert it to []types.Balance
	countmonth := 1
	countday := 0
	initialDate:= c.DateQuantities[0].Date
	prices := *coindata.Prices
	for i, datequantity := range c.DateQuantities {
		for countday < len(prices) {
			dateNow := initialDate.Add(time.Hour * time.Duration(24*countday))
			if i!=len(c.DateQuantities)-1 && (dateNow.After(c.DateQuantities[i+1].Date) || dateNow.Equal(c.DateQuantities[i+1].Date)) {
				break
			}

			balances = append(balances, types.NewBalance(c.CoinType, prices[countday][1],
				dateNow,
				prices[countday][1]*float32(datequantity.Quantity), vsCurrency,datequantity.Quantity))

			countmonth += 1

			if countmonth == 13 {
				countmonth = 1
			}
			if countmonth%2 == 0 {
				countday += 30
			} else {
				countday += 31
			}

		}

	}

	return balances, nil
}

// This get multiple coins and return their total value from Jan 2020 to now 
// It assume the coins in coins has same time period
// It have such a bad complexity :( 
func TotalValueBalanceSheet(coins []types.Coin, vsCurrency string, CG *coingecko.Client)(types.TotalBalances,error){
	coinsbalances := make([]types.Balances,len(coins))
	for i,c := range coins{
		b,err:=ParseBalanceSheet(c,vsCurrency,CG)
		if err!=nil{
			return nil,err
		}
		coinsbalances[i]=b
	}
	
	totalValue :=make(types.TotalBalances,len(coinsbalances[0]))

	for i,b:=range coinsbalances[0]{
		coinTotalValue:=float32(0)
		CoinBalance:=make([]types.Balance,len(coins))
		for j,_ := range coinsbalances{
			coinTotalValue+=coinsbalances[j][i].Balance
			CoinBalance[j]=coinsbalances[j][i]
		}

		totalValue[i]=types.NewTotalBalance(CoinBalance,coinTotalValue,b.Date,)
	}
	return totalValue,nil
}

