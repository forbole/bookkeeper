package coinApi

import (
	"fmt"
	"math/big"
	"net/http"
	"time"

	coingecko "github.com/superoo7/go-gecko/v3"
)

func GetPriceFromCoingecko(coinId, vsCurrency string) (*big.Float, error) {
	// get coin price
	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}

	CG := coingecko.NewClient(httpClient)

	singlePrice, err := CG.CoinsMarket(vsCurrency, []string{coinId}, "", 0, 0, false, nil)
	if err != nil {
		return nil, err
	}
	if len(*singlePrice) == 0 {
		return nil, fmt.Errorf("Error getting coinsmarket")
	}
	coinprice := new(big.Float).SetFloat64((*singlePrice)[0].CurrentPrice)
	fmt.Println(coinprice)
	return coinprice, nil
}
