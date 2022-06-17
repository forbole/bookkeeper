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
	"github.com/rs/zerolog/log"
)

func GetRewardCommission(api *client.SubscanClient, address string, denom types.Denom, vsCurrency string, from int64) (*tabletypes.AddressDateRewardPrice, error) {
	//var rewardPrice tabletypes.DateRewardPriceTable
	log.Trace().Str("module", "subtrate").Msg("GetRewardCommission")

	rewardList, err := GetRewardSlash(api, address, from)
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

		price, err := coinApi.GetCryptoPriceFromDate(timestamp, denom.CoinId, vsCurrency)
		if err != nil {
			return nil, fmt.Errorf("Cannot get crypto price:%s", err)
		}

		rewardPrice := new(big.Float).Mul(reward, price)

		timeRewardPrice[i] = tabletypes.NewDateRewardPriceRow(timestamp, reward, new(big.Float).SetInt64(0),
			denom.CoinId, rewardPrice, new(big.Float).SetInt64(0))

	}
	addressRewardPrice := tabletypes.NewAddressDateRewardPrice(address, timeRewardPrice)
	return &addressRewardPrice, nil
}

func GetRewardSlash(api *client.SubscanClient, address string, from int64) ([]subtratetypes.List, error) {
	log.Trace().Str("module", "subtrate").Msg("GetRewardSlash")

	requestUrl := "/api/scan/account/reward_slash"
	var list []subtratetypes.List

	type Payload struct {
		Row     int    `json:"row"`
		Page    int    `json:"page"`
		Address string `json:"address"`
	}

	count := math.MaxInt
	row := 50

	for page := 0; (page+1)*row < count; page++ {
		payload := Payload{
			Row:     row,
			Page:    page,
			Address: address,
		}
		var rewardSlash subtratetypes.RewardSlash
		err := api.CallApi(requestUrl, payload, &rewardSlash)
		if err != nil {
			return nil, fmt.Errorf("cannot get rewardSlash:%s", err)
		}
		if count == math.MaxInt {
			count = rewardSlash.Data.Count
		}
		rewardList := rewardSlash.Data.List
		if len(rewardList) == 0 {
			break
		}
		list = append(list, rewardList...)
		if int64(rewardList[len(rewardList)-1].BlockTimestamp) < from {
			break
		}
	}

	return list, nil
}
