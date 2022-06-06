package coinApi

import (
	"fmt"
	"math/big"
	"net/http"
	"strings"
	"time"

	"github.com/rs/zerolog/log"
	coingecko "github.com/superoo7/go-gecko/v3"
)

// GetCryptoPrice get cypto price now
func GetCryptoPrice(coinId, vsCurrency string) (*big.Float, error) {
	// Prevent call limit
	time.Sleep(time.Second)
	log.Trace().Str("module", "coinApi").Msg("GetCryptoPrice")

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
	log.Trace().Str("module", "coinApi").Msg("GetCryptoPriceFromDate")
	// Prevent call limit
	time.Sleep(time.Second)

	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}

	cg := coingecko.NewClient(httpClient)
	fmt.Println(date)
	
	prices, err := cg.CoinsIDHistory(coinid, date.Format("02-01-2006"), false)
	if err != nil&& (strings.Contains(err.Error(),"1015")||
			strings.Contains(err.Error(),"context deadline exceeded")){
		log.Error().Str("module", "coinApi").Msg(fmt.Sprintf("Sleep:%s",err.Error()))
		time.Sleep(time.Minute)
		return GetCryptoPriceFromDate(date,coinid,vsCurrency)
	}
	if err!=nil{
		return nil,err
	}

	if prices.MarketData==nil{
		// Set the coin value to 0 if the specific date don't have record
		log.Error().Str("module", "coinApi").Msg(fmt.Sprintf("coingecko don't have record for the date:%s",date))
		return new(big.Float).SetInt64(0),nil
	}

	price := new(big.Float).SetFloat64(prices.MarketData.CurrentPrice[vsCurrency])

	return price, nil

}
