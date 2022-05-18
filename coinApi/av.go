package coinApi

import (
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"os"
	"strings"
	"time"

	cointypes "github.com/forbole/bookkeeper/coinApi/types"
)

// GetCurrencyPrice get currency from AlphaVantage
func GetCurrencyPrice(coinId, vsCurrency string) (*big.Float, error) {
	query := fmt.Sprintf("https://www.alphavantage.co/query?function=CURRENCY_EXCHANGE_RATE&from_currency=%s&to_currency=%s&apikey=%s",
		coinId, vsCurrency, os.Getenv("AV_API_KEY"))
	//fmt.Println(query)
	var bz []byte
	resp, err := http.Get(query)
	if err != nil {
		return nil, fmt.Errorf("Fail to get tx from rpc:%s", err)
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Fail to get tx from rpc:Status :%s", resp.Status)
	}

	bz, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	//retry if api call exceed 5 call per day
	for strings.Contains(string(bz), "Our standard API call frequency is 5 calls per minute and 500 calls per day.") {
		resp.Body.Close()
		//fmt.Println("Exceed limit, sleep")
		time.Sleep(1 * time.Minute)

		//fmt.Println(query)
		resp, err := http.Get(query)

		if err != nil {
			return nil, fmt.Errorf("Fail to get tx from rpc:%s", err)
		}
		if resp.StatusCode != 200 {
			return nil, fmt.Errorf("Fail to get tx from rpc:Status :%s", resp.Status)
		}

		bz, err = io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
	}

	defer resp.Body.Close()

	var exchangeRate cointypes.ExchangeRate
	err = json.Unmarshal(bz, &exchangeRate)
	if err != nil {
		return nil, fmt.Errorf("Fail to marshal:%s", err)
	}

	rate, ok := new(big.Float).SetString(exchangeRate.RealtimeCurrencyExchangeRate.FiveExchangeRate)
	if !ok {
		return nil, fmt.Errorf("Fail to set string:%s", string(bz))
	}
	return rate, nil
}
