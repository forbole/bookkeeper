package tabletypes

import (
	"fmt"
	"math/big"
	"time"
)

// DateRewardPriceRow is a table with date, reward, commission and price on that time
type DateRewardPriceRow struct {
	Date            time.Time
	Reward          *big.Float
	Commission      *big.Float
	Denom           string
	RewardPrice     *big.Float
	CommissionPrice *big.Float
}

func NewDateRewardPriceRow(date time.Time, reward *big.Float, commission *big.Float, denom string,
	rewardPrice *big.Float, commissionPrice *big.Float) DateRewardPriceRow {
	return DateRewardPriceRow{
		Commission:      commission,
		Reward:          reward,
		Denom:           denom,
		Date:            date,
		RewardPrice:     rewardPrice,
		CommissionPrice: commissionPrice,
	}
}

type DateRewardPriceTable []DateRewardPriceRow

type AddressDateRewardPrice struct {
	Address string
	Rows    []DateRewardPriceRow
}

func NewAddressDateRewardPrice(address string, rows []DateRewardPriceRow) AddressDateRewardPrice {
	return AddressDateRewardPrice{
		Address: address,
		Rows:    rows,
	}
}

func (v AddressDateRewardPrice) GetCSV() string {
	csv := "Date,commission,reward,unit,commission_$value,reward_$value\n"
	for _, row := range v.Rows {
		csv += fmt.Sprintf("%s,%f,%f,%s,%f,%f\n", row.Date, row.Commission, row.Reward, row.Denom, row.CommissionPrice, row.RewardPrice)
	}

	return csv
}

// GetDelegationCSV when the source use delegation instead of commission
// show delegation instead of commission
func (v AddressDateRewardPrice) GetDelegationCSV() string {
	csv:="Date,delegation,$delegation_value\n"
	for _, row := range v.Rows {
		csv += fmt.Sprintf("%s,%f,%s,%f\n", row.Date, row.Reward, row.Denom, row.RewardPrice)
	}

	return csv
}