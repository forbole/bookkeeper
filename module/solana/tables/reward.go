package tables

import (
	"fmt"
	"math"
	"math/big"
	"time"

	"github.com/forbole/bookkeeper/coinApi"
	"github.com/forbole/bookkeeper/module/solana/client"
	"github.com/forbole/bookkeeper/module/solana/utils"
	"github.com/forbole/bookkeeper/types"

	"github.com/forbole/bookkeeper/types/tabletypes"
)

func GetStakeRewardForPubKey(solana types.Solana, from int64, vsCurrency string, client *client.SolanaBeachClient) ([]tabletypes.AddressDateRewardPrice, error) {
	var addressRewardPrice []tabletypes.AddressDateRewardPrice
	addresses, err := utils.GetSelfDelegatorAddresses(solana.PubKey, solana.ValidatorIdentity, client)
	if err != nil {
		return nil, err
	}

	epoch, err := utils.GetEpochByTime(from, client)
	if err != nil {
		return nil, err
	}

	for _, address := range addresses {
		table, err := GetRewardFromAddress(address, solana.Denom, epoch, vsCurrency, client)
		if err != nil {
			return nil, err
		}
		addressRewardPrice = append(addressRewardPrice, *table)

	}

	return addressRewardPrice, nil
}

func GetRewardFromAddress(address string, denom types.Denom, epoch int, vsCurrency string, client *client.SolanaBeachClient) (*tabletypes.AddressDateRewardPrice, error) {
	reward, err := client.GetStakeReward(address, epoch)
	if err != nil {
		return nil, err
	}
	var dateRewardPrice []tabletypes.DateRewardPriceRow
	for _, r := range reward {
		timestamp := time.Unix(int64(r.Timestamp), 0)
		rewardRaw := new(big.Float).SetInt64(int64(r.Amount))
		exp := new(big.Float).SetFloat64(math.Pow10(-1 * denom.Exponent))
		reward := new(big.Float).Mul(rewardRaw, exp)

		price, err := coinApi.GetCryptoPriceFromDate(timestamp, denom.CoinId, vsCurrency)
		if err != nil {
			return nil, err
		}
		fmt.Println(price)

		rewardPrice := new(big.Float).Mul(reward, price)

		dateRewardPrice = append(dateRewardPrice, tabletypes.NewDateRewardPriceRow(timestamp, reward,
			new(big.Float).SetInt64(0), denom.Denom, rewardPrice, new(big.Float).SetInt64(0)))
	}
	addressRewardPrice := tabletypes.NewAddressDateRewardPrice(address, dateRewardPrice)

	return &addressRewardPrice, nil
}
