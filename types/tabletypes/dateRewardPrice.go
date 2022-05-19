package tabletypes

import (
	"fmt"
	"math/big"
	"time"
)

// DateRewardPriceRow is a table with date, reward, commission and price on that time
type DateRewardPriceRow struct{
	Date time.Time
	Reward     *big.Int
	Commission *big.Int
	Denom      string
	RewardPrice     *big.Float
	CommissionPrice *big.Float
}

func NewDateRewardPriceRow(date time.Time,reward *big.Int,commission *big.Int,denom string,
	rewardPrice *big.Float,commissionPrice *big.Float)DateRewardPriceRow{
	return DateRewardPriceRow{
		Commission: commission,
		Reward:     reward,
		Denom:      denom,
		Date:date,
		RewardPrice:     rewardPrice,
		CommissionPrice:      commissionPrice,
	}
} 

type DateRewardPriceTable []DateRewardPriceRow

type AddressDateRewardPrice struct{
	Address string
	Rows []DateRewardPriceRow
}

func NewAddressDateRewardPrice(address string,rows []DateRewardPriceRow)AddressDateRewardPrice{
	return AddressDateRewardPrice{
		Address: address,
		Rows:rows,
	}
}

func (v DateRewardPriceTable)GetCSV()string{
	csv:="Date,commission,reward,denom,commission_price,reward_price\n"
	for _,row:=range v{
		csv+=fmt.Sprintf("%s,%f,%f,%s,%f,%f",row.Date,row.Commission,row.Reward,row.Denom,row.CommissionPrice,row.RewardPrice)
	}
	return csv
}