package tables

import (
	"fmt"
	"math"
	"math/big"
	"time"

	"github.com/forbole/bookkeeper/coinApi"
	elrondClient "github.com/forbole/bookkeeper/module/elrond/client"
	elrondtypes "github.com/forbole/bookkeeper/module/elrond/types"
	"github.com/forbole/bookkeeper/types"

	"github.com/forbole/bookkeeper/types/tabletypes"
	"github.com/rs/zerolog/log"
)

// GetTxs
func GetTxs(client *elrondClient.ElrondClient, address, contract string, from int64, vsCurrency string) ([]elrondtypes.TxResult, error) {
	log.Trace().Str("module", "elrond").Msg("GetTxs")

	txs, err := client.GetTxs(address, "")
	if err != nil {
		return nil, err
	}
	fmt.Println(len(txs))

	var results []elrondtypes.TxResult
	for _, tx := range txs {
		if tx.Status == "fail" {
			continue
		}
		txResult, err := client.GetTxResult(tx.TxHash)
		if err != nil {
			return nil, err
		}
		results = append(results, *txResult)
	}
	return results, nil

}

func GetValueRow(txResult []elrondtypes.TxResult, vsCurrency string, denom types.Denom, validatorContract string) (tabletypes.DateValueRows, error) {
	log.Trace().Str("module", "elrond").Msg("GetValueRow")

	exp := new(big.Float).SetFloat64(math.Pow10(-1 * denom.Exponent))
	d := denom.CoinId

	var dateValueRows []tabletypes.DateValueRow
	for _, result := range txResult {
		timestamp := time.Unix(int64(result.Timestamp), 0)
		price, err := coinApi.GetCryptoPriceFromDate(timestamp, d, vsCurrency)
		if err != nil {
			return nil, err
		}

		for _, r := range result.Results {
			if !(r.Receiver == validatorContract || r.Sender == validatorContract) {
				continue
			}
			rewardRaw, ok := new(big.Float).SetString(r.Value)
			if !ok {
				return nil, fmt.Errorf("Cannot convert to int:%s", r.Value)
			}

			reward := new(big.Float).Mul(rewardRaw, exp)
			convertPrice := new(big.Float).Mul(price, reward)
			dateValueRow := tabletypes.NewDateValueRow(timestamp, reward, d, convertPrice, result.TxHash, result.Function)
			dateValueRows = append(dateValueRows, dateValueRow)
		}

	}

	return dateValueRows, nil

}
