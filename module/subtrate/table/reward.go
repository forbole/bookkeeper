package table

import (
	"fmt"
	"math"
	"math/big"
	"time"

	"github.com/forbole/bookkeeper/coinApi"
	"github.com/forbole/bookkeeper/module/subtrate/client"
	subtratetypes "github.com/forbole/bookkeeper/module/subtrate/types"
	"github.com/forbole/bookkeeper/types"
	tabletypes "github.com/forbole/bookkeeper/types/tabletypes"
)

func GetRewardCommission(api *client.SubscanClient, address string, denom types.Denom, vsCurrency string) (*tabletypes.AddressDateRewardPrice, error) {
	//var rewardPrice tabletypes.DateRewardPriceTable
	rewardList, err := GetRewardSlash(api, address)
	if err != nil {
		return nil, err
	}

	timeRewardPrice := make(tabletypes.DateRewardPriceTable, len(rewardList))
	for i, list := range rewardList {
		timestamp := time.Unix(int64(list.BlockTimestamp), 0)

		exponent := new(big.Float).SetFloat64((math.Pow(10, float64(-1*denom.Exponent))))
		amount, ok := new(big.Float).SetString(list.Amount)
		if !ok {
			return nil, fmt.Errorf("Cannot convert amount to big.Float")
		}
		reward := new(big.Float).Mul(amount, exponent)

		price, err := coinApi.GetCryptoPriceFromDate(timestamp,denom.CoinId, vsCurrency)
		if err != nil {
			return nil, err
		}

		rewardPrice := new(big.Float).Mul(reward, price)

		timeRewardPrice[i] = tabletypes.NewDateRewardPriceRow(timestamp, reward, new(big.Float).SetInt64(0),
			denom.CoinId, rewardPrice, new(big.Float).SetInt64(0))
	}
	addressRewardPrice:=tabletypes.NewAddressDateRewardPrice(address,timeRewardPrice)
	return &addressRewardPrice, nil
}

func GetRewardSlash(api *client.SubscanClient, address string) ([]subtratetypes.List, error) {
	requestUrl := "/api/v2/scan/account/reward_slash"
	var list []subtratetypes.List

	type Payload struct {
		Row     int    `json:"row"`
		Page    int    `json:"page"`
		Address string `json:"address"`
	}

	page := 0
	count := math.MaxInt
	row := 20

	for ; (page+1)*row < count; page++ {
		payload := Payload{
			Row:     row,
			Page:    page,
			Address: address,
		}
		var rewardSlash subtratetypes.RewardSlash
		err := api.CallApi(requestUrl, payload, &rewardSlash)
		if err != nil {
			return nil, err
		}
		if count == math.MaxInt {
			count = rewardSlash.Data.Count
		}
		list = append(list, rewardSlash.Data.List...)
	}

	return list, nil
}
