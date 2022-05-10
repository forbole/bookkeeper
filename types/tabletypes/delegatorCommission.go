package tabletypes

import (
	"fmt"
	"math/big"
)

type RewardCommission struct {
	TxHash     string
	Height     int
	Commission *big.Int
	Reward     *big.Int
	Denom      string
}

func NewRewardCommission(txHash string, height int, denom string, commission, reward *big.Int) RewardCommission {
	return RewardCommission{
		TxHash:     txHash,
		Height:     height,
		Commission: commission,
		Reward:     reward,
		Denom:      denom,
	}
}

type RewardCommissions []RewardCommission

type AddressRewardCommission struct {
	Address string
	Rows    RewardCommissions
}

func NewAddressRewardCommission(address string, rewardCommissions RewardCommissions) AddressRewardCommission {
	return AddressRewardCommission{
		Address: address,
		Rows:    rewardCommissions,
	}
}

func (v RewardCommissions) GetCSV() string {
	outputcsv := "height,txHash,Commission,Delegator_Reward,denom\n"
	commissionSum := big.NewInt(0)
	rewardSum := big.NewInt(0)
	for _, b := range v {
		outputcsv += fmt.Sprintf("%d,%s,%v,%v,%s\n",
			b.Height, b.TxHash, b.Commission, b.Reward, b.Denom)
		commissionSum.Add(commissionSum, b.Commission)
		rewardSum.Add(rewardSum, b.Reward)
	}
	//outputcsv+=fmt.Sprintf("\n Sum, ,%v,%v\n",commissionSum,rewardSum)

	return outputcsv
}
