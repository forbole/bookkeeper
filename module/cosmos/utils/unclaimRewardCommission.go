package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"time"

	cosmostypes "github.com/forbole/bookkeeper/module/cosmos/types"
	"github.com/forbole/bookkeeper/types/tabletypes"

	"github.com/rs/zerolog/log"
)

// GetUnclaimedRewardCommission get unclaimed commission and reward from a validator address
func GetUnclaimedRewardCommission(lcd string, address string) ([]tabletypes.MonthyReportRow, error) {
	var monthyReportRows []tabletypes.MonthyReportRow
	now := time.Now()
	commission, err := getUnclaimCommission(lcd, address)
	if err != nil {
		return nil, err
	}

	reward, err := getUnclaimReward(lcd, address)
	if err != nil {
		return nil, err
	}

	for _, c := range commission {
		unclaimedCommission, ok := new(big.Float).SetString(c.Amount)
		if !ok {
			return nil, fmt.Errorf("Cannot read unclaimecd Commission:%s", c.Amount)
		}
		// find the corresponding reward
		unclaimedReward := new(big.Float).SetInt64(0)
		for _, r := range reward {
			if r.Denom == c.Denom {
				newReward, ok := new(big.Float).SetString(r.Amount)
				if !ok {
					return nil, fmt.Errorf("Cannot read unclaimecd Reward:%s", r.Amount)
				}
				unclaimedReward = newReward
			}
		}
		monthyReportRows = append(monthyReportRows,
			tabletypes.NewMonthyReportRow(now, now, unclaimedCommission, unclaimedReward, c.Denom))
	}
	return monthyReportRows, nil
}

// getUnclaimCommission get unclaimed commission from a validator address
func getUnclaimCommission(lcd string, address string) ([]cosmostypes.DenomAmount, error) {
	query := fmt.Sprintf(`%s/cosmos/distribution/v1beta1/validators/%s/commission`,
		lcd, address)
	log.Trace().Str("module", "cosmos").Str("query", query).Msg("get unclaim commission")
	resp, err := http.Get(query)
	if err != nil {
		return nil, fmt.Errorf("Fail to get tx from rpc:%s", err)
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Fail to get tx from rpc:Status :%s", resp.Status)
	}

	defer resp.Body.Close()

	bz, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var txSearchRes cosmostypes.Commission
	err = json.Unmarshal(bz, &txSearchRes)
	if err != nil {
		return nil, fmt.Errorf("Fail to marshal:%s", err)
	}
	return txSearchRes.Commission.Commission, nil
}

// getUnclaimReward get unclaimed reward from a validator address
func getUnclaimReward(lcd string, address string) ([]cosmostypes.DenomAmount, error) {
	query := fmt.Sprintf(`%s/cosmos/distribution/v1beta1/validators/%s/outstanding_rewards`,
		lcd, address)
	log.Trace().Str("module", "cosmos").Str("query", query).Msg("get unclaim reward")
	resp, err := http.Get(query)
	if err != nil {
		return nil, fmt.Errorf("Fail to get tx from rpc:%s", err)
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Fail to get tx from rpc:Status :%s", resp.Status)
	}

	defer resp.Body.Close()

	bz, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var txSearchRes cosmostypes.Rewards
	err = json.Unmarshal(bz, &txSearchRes)
	if err != nil {
		return nil, fmt.Errorf("Fail to marshal:%s", err)
	}
	return txSearchRes.Rewards.Rewards, nil
}
