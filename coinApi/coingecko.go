package coinApi

import (
	"fmt"
	"math/big"
	"net/http"
	"time"

	coingecko "github.com/superoo7/go-gecko/v3"
)

// GetCryptoPrice get cypto price now
func GetCryptoPrice(coinId, vsCurrency string) (*big.Float, error) {
	// get coin price
	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}

	cg := coingecko.NewClient(httpClient)

	singlePrice, err := cg.CoinsMarket(vsCurrency, []string{coinId}, "", 0, 0, false, nil)
	if err != nil {
		return nil, err
	}
	if len(*singlePrice) == 0 {
		return nil, fmt.Errorf("Error getting coinsmarket")
	}
	coinprice := new(big.Float).SetFloat64((*singlePrice)[0].CurrentPrice)
	//fmt.Println(coinprice)
	return coinprice, nil
}

// GetPriceFromDate get crypto price by required date
func GetCryptoPriceFromDate(date time.Time, coinid, vsCurrency string) (*big.Float, error) {
	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}

	cg := coingecko.NewClient(httpClient)

	prices, err := cg.CoinsIDHistory(coinid, date.Format("02-01-2006"), false)
	if err != nil {
		return nil, err
	}
	price := new(big.Float).SetFloat64(prices.MarketData.CurrentPrice[vsCurrency])

	return price, nil

}
