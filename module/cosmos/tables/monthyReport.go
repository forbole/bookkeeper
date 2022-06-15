package tables

import (
	"math/big"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/forbole/bookkeeper/module/cosmos/utils"
	types "github.com/forbole/bookkeeper/types"
	tabletypes "github.com/forbole/bookkeeper/types/tabletypes"
)

// GetMonthyReport get monthy report between certain period of time
func GetMonthyReport(details types.CosmosDetails, period types.Period) ([]tabletypes.AddressMonthyReport, error) {
	log.Trace().Str("module", "cosmos").Msg("GetMonthyReport")

	var monthyReports []tabletypes.AddressMonthyReport

	from := time.Unix(period.From, 0)

	to := time.Unix(period.To, 0)

	for _, validator := range details.Validators {
		validatorReport, err := GetMonthyReportForValidator(validator, from, to, details.LcdEndpoint, details.RpcEndpoint)
		if err != nil {
			return nil, err
		}

		monthyReports = append(monthyReports, tabletypes.NewAddressMonthyReport(validator.ValidatorAddress, validatorReport))
	}

	for _, address := range details.FundHoldingAccount {
		report, err := GetMonthyReportForSingleAddress(address, from, to, details.LcdEndpoint, details.RpcEndpoint)
		if err != nil {
			return nil, err
		}

		monthyReports = append(monthyReports, tabletypes.NewAddressMonthyReport(address, report))
	}
	return monthyReports, nil
}

// GetMonthyReportForSingleAddress get monthy report for a single validator and its delegator address
func GetMonthyReportForSingleAddress(address string, from time.Time, to time.Time, lcd, rpc string) ([]tabletypes.MonthyReportRow, error) {
	log.Trace().Str("module", "cosmos").Msg("GetMonthyReportForSingleAddress")

	targetHeight, err := utils.GetHeightByDate(from, lcd)
	if err != nil {
		return nil, err
	}

	txs, err := GetTxsForAnAddress(address, rpc, targetHeight)
	if err != nil {
		return nil, err
	}

	rewardCommission, err := GetRewardCommission(*txs)
	if err != nil {
		return nil, err
	}

	monthyReportRowsFromRewardCommission, err := GetMonthyReportFromRewardCommission(rewardCommission.Rows, to, from, lcd)
	if err != nil {
		return nil, err
	}

	return monthyReportRowsFromRewardCommission, nil
}

// GetMonthyReportForValidator get monthy report for a validator including unclaimed reward
func GetMonthyReportForValidator(validatorDetail types.ValidatorDetail, from time.Time, to time.Time, lcd, rpc string) ([]tabletypes.MonthyReportRow, error) {
	log.Trace().Str("module", "cosmos").Msg("GetMonthyReportForValidator")

	
	unclaimedRewardCommission, err := utils.GetUnclaimedRewardCommission(lcd, validatorDetail.ValidatorAddress)
	if err != nil {
		return nil, err
	}

	report, err := GetMonthyReportForSingleAddress(validatorDetail.SelfDelegationAddress, from, to, lcd, rpc)
	if err != nil {
		return nil, err
	}

	validatorMonthyReport := append(unclaimedRewardCommission, report...)
	return validatorMonthyReport, nil
}

// It pass the RewardCommission and output a monthy report
func GetMonthyReportFromRewardCommission(rows tabletypes.RewardCommissions, to time.Time, from time.Time, lcdEndpoint string) ([]tabletypes.MonthyReportRow, error) {
	log.Trace().Str("module", "cosmos").Msg("GetMonthyReportFromRewardCommission")

	
	t := to
	i := 0
	var monthyReportRows []tabletypes.MonthyReportRow
	for t.After(from) && len(rows) > i {
		targetHeight, err := utils.GetHeightByDate(t, lcdEndpoint)
		if err != nil {
			return nil, err
		}

		if rows[i].Height < targetHeight {
			monthyReportRows = append(monthyReportRows,
				tabletypes.NewMonthyReportRow(t, to, new(big.Float).SetFloat64(0),
					new(big.Float).SetFloat64(0), rows[i].Denom))
			t = *(utils.LastMonth(t))
			continue
		}
		// A record for a month include different denom. So this struct should be denom vs
		//the reward and commission value in this denom
		recordForMonth := make(map[string]*RewardCommission)
		for ; len(rows) > i && rows[i].Height > targetHeight; i++ {
			// recordForMonth have denom as key and sum up reward and commission for the month
			denomEntry, ok := recordForMonth[rows[i].Denom]
			if !ok {
				e := &RewardCommission{
					Commission: big.NewFloat(0),
					Reward:     big.NewFloat(0),
				}
				denomEntry = e
			}

			c := new(big.Float).SetInt(rows[i].Commission)
			r := new(big.Float).SetInt(rows[i].Reward)

			commission := new(big.Float).Add(denomEntry.Commission, c)
			reward := new(big.Float).Add(denomEntry.Reward, r)
			denomEntry.Commission = new(big.Float).Set(commission)
			denomEntry.Reward = new(big.Float).Set(reward)
			recordForMonth[rows[i].Denom] = denomEntry
			//fmt.Println(denomEntry.Commission)
		}

		for key, element := range recordForMonth {
			to := *(utils.NextMonth(t))
			if time.Now().Before(*(utils.NextMonth(t))) {
				to = time.Now()
			}
			//fmt.Println("Key:", key, "=>", "Element:", element)
			//fmt.Println(t)
			//fmt.Println(to)

			monthyReportRows = append(monthyReportRows,
				tabletypes.NewMonthyReportRow(t, to, element.Commission, element.Reward, key))
		}
		t = *(utils.LastMonth(t))
	}
	return monthyReportRows, nil
}

type RewardCommission struct {
	Commission *big.Float
	Reward     *big.Float
}
