package tables

import (
	"fmt"
	"math"
	"math/big"
	"time"

	"github.com/forbole/bookkeeper/coinApi"
	"github.com/forbole/bookkeeper/module/cosmos/utils"
	types "github.com/forbole/bookkeeper/types"
	tabletypes "github.com/forbole/bookkeeper/types/tabletypes"
)

// DateRewardValue table is the table include converted value reward and commission on that day
func GetDateRewardValueFromDetails(details types.CosmosDetails, period types.Period, vsCurrency string) (
	[]tabletypes.AddressDateRewardPrice, error) {
	addressDateRewardPrice := make([]tabletypes.AddressDateRewardPrice, len(details.Validators)+len(details.FundHoldingAccount))
	date := time.Unix(period.From, 0)
	targetHeight, err := utils.GetHeightByDate(date, details.LcdEndpoint)
	if err != nil {
		return nil, err
	}
	i := 0
	denomMap := ConvertDenomToMap(details.Denom)
	for _, validator := range details.Validators {
		dateRewardValue, err := DateRewardCommissionValueForAnAddress(validator.SelfDelegationAddress,
			details.LcdEndpoint, details.RpcEndpoint, targetHeight, denomMap, vsCurrency)
		if err != nil {
			return nil, err
		}
		addressDateRewardPrice[i] = tabletypes.NewAddressDateRewardPrice(validator.ValidatorAddress, dateRewardValue)
		i++
	}
	for _, address := range details.FundHoldingAccount {
		dateRewardValue, err := DateRewardCommissionValueForAnAddress(address,
			details.LcdEndpoint, details.RpcEndpoint, targetHeight, denomMap, vsCurrency)
		if err != nil {
			return nil, err
		}
		addressDateRewardPrice[i] = tabletypes.NewAddressDateRewardPrice(address, dateRewardValue)
		i++
	}
	return addressDateRewardPrice, nil
}

func DateRewardCommissionValueForAnAddress(address string, lcd string, rpc string, targetHeight int, denoms denomMap, vsCurrency string) (
	[]tabletypes.DateRewardPriceRow, error) {
	txs, err := GetTxsForAnAddress(address, rpc, targetHeight)
	if err != nil {
		return nil, err
	}

	rewardCommission, err := GetRewardCommission(*txs)
	if err != nil {
		return nil, err
	}

	dateRewardValues, err := GetDateRewardCommissionValue(rewardCommission.Rows, denoms, vsCurrency, lcd)
	if err != nil {
		return nil, err
	}

	return dateRewardValues, nil

}
func GetDateRewardCommissionValue(v tabletypes.RewardCommissions, denomMap denomMap, vsCurrency string, lcd string) (
	[]tabletypes.DateRewardPriceRow, error) {
	DateRewardPriceRow := make([]tabletypes.DateRewardPriceRow, len(v))
	for i, r := range v {
		date, err := utils.GetTimeByHeight(r.Height, lcd)
		if err != nil {
			return nil, err
		}

		// If that is not in the denom list, not getting price
		if _, ok := denomMap[r.Denom]; !ok {
			DateRewardPriceRow[i] = tabletypes.NewDateRewardPriceRow(*date, new(big.Float).SetInt(r.Reward), new(big.Float).SetInt(r.Commission), r.Denom, new(big.Float).SetInt64(0), new(big.Float).SetInt64(0))
			continue
		}

		var price *big.Float
		if denomMap[r.Denom].Cointype == "crypto" {
			price, err = coinApi.GetCryptoPriceFromDate(*date, denomMap[r.Denom].CoinId, vsCurrency)
			if err != nil {
				return nil, err
			}
		} else if denomMap[r.Denom].Cointype == "stablecoin" {
			price, err = coinApi.GetCurrencyPrice(denomMap[r.Denom].CoinId, vsCurrency)
			if err != nil {
				return nil, err
			}
		}
		fmt.Println(r.Denom)
		commission := new(big.Float).Mul(new(big.Float).SetInt(r.Commission), denomMap[r.Denom].Exponent)
		reward := new(big.Float).Mul(new(big.Float).SetInt(r.Reward), denomMap[r.Denom].Exponent)

		commissionPrice := new(big.Float).Mul(commission, price)
		rewardPrice := new(big.Float).Mul(reward, price)

		DateRewardPriceRow[i] = tabletypes.NewDateRewardPriceRow(*date, reward, commission, denomMap[r.Denom].CoinId, rewardPrice, commissionPrice)

	}
	return DateRewardPriceRow, nil
}

type denomDetails struct {
	Exponent *big.Float `json:"exponent"`
	CoinId   string     `json:"coin_id"`
	Cointype string     `json:"cointype"`
}

type denomMap map[string]denomDetails

func ConvertDenomToMap(denoms []types.Denom) denomMap {
	denomDetailsMap := make(map[string]denomDetails)
	for _, d := range denoms {
		exponent := new(big.Float).SetFloat64((math.Pow(10, float64(-1*d.Exponent))))

		denomDetailsMap[d.Denom] = denomDetails{
			Exponent: exponent,
			CoinId:   d.CoinId,
			Cointype: d.Cointype,
		}
	}
	return denomDetailsMap
}
