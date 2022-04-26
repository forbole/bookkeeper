package coinApi

import (
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"os"

	cointypes "github.com/forbole/bookkeeper/coinApi/types"
)

// GetPriceFromAV get currency from AlphaVantage
func GetPriceFromAV(coinId,vsCurrency string)(*big.Float,error){
	query:=fmt.Sprintf("https://www.alphavantage.co/query?function=CURRENCY_EXCHANGE_RATE&from_currency=%s&to_currency=%s&apikey=%s",
			coinId,vsCurrency,os.Getenv("AV_API_KEY"))
			fmt.Println(query)
			resp, err := http.Get(query)
			if err != nil {
				return nil, fmt.Errorf("Fail to get tx from rpc:%s", err)
			}
			if resp.StatusCode != 200 {
				return nil, fmt.Errorf("Fail to get tx from rpc:Status :%s", resp.Status)
			}

			defer resp.Body.Close()

			bz, err := io.ReadAll(resp.Body)

			var exchangeRate cointypes.ExchangeRate
			err = json.Unmarshal(bz, &exchangeRate)
			if err != nil {
				return nil, fmt.Errorf("Fail to marshal:%s", err)
			}

			rate,ok:=new(big.Float).SetString(exchangeRate.RealtimeCurrencyExchangeRate.FiveExchangeRate)
			if !ok{
				return nil,fmt.Errorf("Fail to set string:%s",exchangeRate.RealtimeCurrencyExchangeRate.FiveExchangeRate)
			}
			return rate,nil
}

