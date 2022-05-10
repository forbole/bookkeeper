package tables

import (
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"time"

	cosmostypes "github.com/forbole/bookkeeper/module/cosmos/types"
	"github.com/forbole/bookkeeper/module/cosmos/utils"
	types "github.com/forbole/bookkeeper/types"
	tabletypes "github.com/forbole/bookkeeper/types/tabletypes"
)

// GetMonthyReport get monthy report between certain period of time
func GetMonthyReport(details types.IndividualChain, period types.Period) ([]tabletypes.AddressMonthyReport, error) {
	//var monthyReports []tabletypes.AddressMonthyReport
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
		monthyReportRows, err := getUnclaimedRewardCommission(details.LcdEndpoint, details.Validators[0].ValidatorAddress)
		if err != nil {
			return nil, err
		}

		rewardCommission, err := GetRewardCommission(b)
		if err != nil {
			return nil, err
		}
		rows := rewardCommission.Rows
		fmt.Println(rewardCommission.Rows.GetCSV())

		i := 0
		for t.After(from) && len(rows) > i {
			targetHeight, err := utils.GetHeightByDate(t, details.LcdEndpoint)
			if err != nil {
				return nil, err
			}

			if rows[i].Height < targetHeight {
				monthyReportRows = append(monthyReportRows,
					tabletypes.NewMonthyReportRow(t, to, new(big.Float).SetFloat64(0),
						new(big.Float).SetFloat64(0), rows[i].Denom))
				t = *(lastMonth(t))
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
				fmt.Println(denomEntry.Commission)
			}

			for key, element := range recordForMonth {
				to := *(nextMonth(t))
				if time.Now().Before(*(nextMonth(t))) {
					to = time.Now()
				}
				fmt.Println("Key:", key, "=>", "Element:", element)
				fmt.Println(t)
				fmt.Println(to)

				monthyReportRows = append(monthyReportRows,
					tabletypes.NewMonthyReportRow(t, to, element.Commission, element.Reward, key))
			}
			t = *(lastMonth(t))

		}
		//monthyReports = append(monthyReports, tabletypes.NewAddressMonthyReport(b.Address, monthyReportRows))
		monthyReports[j] = tabletypes.NewAddressMonthyReport(b.Address, monthyReportRows)
	}
	return monthyReports, nil
}

func nextMonth(t time.Time) *time.Time {
	date := t

	day := t.Day()
	if t.Day() > 28 {
		day = 28
	}

	switch t.Month() {
	case time.January:
		date = time.Date(t.Year(), time.February, day, t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())
		break
	case time.February:
		date = time.Date(t.Year(), time.March, day, t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())
		break
	case time.March:
		date = time.Date(t.Year(), time.April, day, t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())
		break
	case time.April:
		date = time.Date(t.Year(), time.May, day, t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())
		break
	case time.May:
		date = time.Date(t.Year(), time.June, day, t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())
		break
	case time.June:
		date = time.Date(t.Year(), time.July, day, t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())
		break
	case time.July:
		date = time.Date(t.Year(), time.August, day, t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())
		break
	case time.August:
		date = time.Date(t.Year(), time.September, day, t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())
		break
	case time.September:
		date = time.Date(t.Year(), time.October, day, t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())
		break
	case time.October:
		date = time.Date(t.Year(), time.November, day, t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())
		break
	case time.November:
		date = time.Date(t.Year(), time.December, day, t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())
		break
	case time.December:
		date = time.Date(t.Year()+1, time.January, day, t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())
		break
	}

	return &date
}

func lastMonth(t time.Time) *time.Time {
	date := t

	day := t.Day()
	if t.Day() > 28 {
		day = 28
	}

	switch t.Month() {
	case time.January:
		date = time.Date(t.Year()-1, time.December, day, t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())
		break
	case time.February:
		date = time.Date(t.Year(), time.January, day, t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())
		break
	case time.March:
		date = time.Date(t.Year(), time.February, day, t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())
		break
	case time.April:
		date = time.Date(t.Year(), time.March, day, t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())
		break
	case time.May:
		date = time.Date(t.Year(), time.April, day, t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())
		break
	case time.June:
		date = time.Date(t.Year(), time.May, day, t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())
		break
	case time.July:
		date = time.Date(t.Year(), time.June, day, t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())
		break
	case time.August:
		date = time.Date(t.Year(), time.July, day, t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())
		break
	case time.September:
		date = time.Date(t.Year(), time.August, day, t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())
		break
	case time.October:
		date = time.Date(t.Year(), time.September, day, t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())
		break
	case time.November:
		date = time.Date(t.Year(), time.October, day, t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())
		break
	case time.December:
		date = time.Date(t.Year(), time.November, day, t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())
		break
	}

	return &date
}

type RewardCommission struct {
	Commission *big.Float
	Reward     *big.Float
}

func getUnclaimedRewardCommission(lcd string, address string) ([]tabletypes.MonthyReportRow, error) {
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

func getUnclaimCommission(lcd string, address string) ([]cosmostypes.DenomAmount, error) {
	query := fmt.Sprintf(`%s/cosmos/distribution/v1beta1/validators/%s/commission`,
		lcd, address)
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

	var txSearchRes cosmostypes.Commission
	err = json.Unmarshal(bz, &txSearchRes)
	if err != nil {
		return nil, fmt.Errorf("Fail to marshal:%s", err)
	}
	return txSearchRes.Commission.Commission, nil
}

func getUnclaimReward(lcd string, address string) ([]cosmostypes.DenomAmount, error) {
	query := fmt.Sprintf(`%s/cosmos/distribution/v1beta1/validators/%s/outstanding_rewards`,
		lcd, address)
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

	var txSearchRes cosmostypes.Rewards
	err = json.Unmarshal(bz, &txSearchRes)
	if err != nil {
		return nil, fmt.Errorf("Fail to marshal:%s", err)
	}
	return txSearchRes.Rewards.Rewards, nil
}
