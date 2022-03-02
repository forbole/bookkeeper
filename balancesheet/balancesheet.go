package balancesheet

import (
	"strconv"
	"time"

	"github.com/HarleyAppleChoi/bookeeper/types"
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
			if i!=len(c.DateQuantities)-1 && dateNow.After(c.DateQuantities[i+1].Date) {
				break
			}

			balances = append(balances, types.NewBalance(&c, prices[countday][1],
				dateNow,
				prices[countday][1]*float32(datequantity.Quantity), vsCurrency))

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
