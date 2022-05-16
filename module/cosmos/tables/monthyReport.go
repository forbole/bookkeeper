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
func GetMonthyReport(details types.IndividualChain, period types.Period) ([]tabletypes.AddressMonthyReport, error) {
	//var monthyReports []tabletypes.AddressMonthyReport
	log.Trace().Str("module", "cosmos").Msg("get monthy report")

	monthyReports := make([]tabletypes.AddressMonthyReport, len(details.FundHoldingAccount))

	from := time.Unix(period.From, 0)

	to := time.Unix(period.To, 0)

	balanceEntries, err := GetTxs(details, period.From)
	if err != nil {
		return nil, err
	}
	if balanceEntries == nil {
		return nil, nil
	}
	t := to
	for j, b := range balanceEntries {
		var monthyReportRows []tabletypes.MonthyReportRow
		monthyReportRows, err := utils.GetUnclaimedRewardCommission(details.LcdEndpoint, details.Validators[0].ValidatorAddress)
		if err != nil {
			return nil, err
		}

		rewardCommission, err := GetRewardCommission(b)
		if err != nil {
			return nil, err
		}
		rows := rewardCommission.Rows
		//fmt.Println(rewardCommission.Rows.GetCSV())

		monthyReportRowsFromRewardCommission,err:=GetMonthyReportForAnAddress(rows,t,from,details.LcdEndpoint)
		if err!=nil{
			return nil,err
		}
		
		monthyReportRows=append(monthyReportRows, monthyReportRowsFromRewardCommission...)
		//monthyReports = append(monthyReports, tabletypes.NewAddressMonthyReport(b.Address, monthyReportRows))
		monthyReports[j] = tabletypes.NewAddressMonthyReport(b.Address, monthyReportRows)
	}
	return monthyReports, nil
}

// It pass the RewardCommission and output a monthy report
func GetMonthyReportForAnAddress(rows tabletypes.RewardCommissions,to time.Time,from time.Time,lcdEndpoint string)([]tabletypes.MonthyReportRow,error){
	t:=to
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
		return monthyReportRows,nil
}

type RewardCommission struct {
	Commission *big.Float
	Reward     *big.Float
}