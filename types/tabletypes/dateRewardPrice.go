package tabletypes

import (
	"math/big"
	"time"
)

// DateRewardCommissionPrice is a table with date, reward, commission and price on that time
type DateRewardCommissionPrice struct{
	Date time.Time
	Reward     *big.Int
	Commission *big.Int
	Denom      string
	RewardPrice     *big.Float
	CommissionPrice *big.Float
}

func NewDateRewardCommissionPrice(date time.Time,reward *big.Int,commission *big.Int,denom string,
	rewardPrice *big.Float,commissionPrice *big.Float)DateRewardCommissionPrice{
	return DateRewardCommissionPrice{
		Commission: commission,
		Reward:     reward,
		Denom:      denom,
		Date:date,
		RewardPrice:     rewardPrice,
		CommissionPrice:      commissionPrice,
	}
} 
