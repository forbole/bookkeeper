package types

import "fmt"

type RewardCommission struct{
	TxHash string
	Height int
	Commission int
	Reward int
	Denom string
}

func NewRewardCommission(txHash string,height int,denom string,commission,reward int)RewardCommission{
	return RewardCommission{
		TxHash: txHash,
		Height:height,
		Commission: commission,
		Reward: reward,
		Denom: denom,
	}
}

type RewardCommissions []RewardCommission

type AddressRewardCommission struct{
	Address string
	Rows RewardCommissions
}

func NewAddressRewardCommission(address string,rewardCommissions RewardCommissions)AddressRewardCommission{
	return AddressRewardCommission{
		Address: address,
		Rows: rewardCommissions,
	}
}

func (v RewardCommissions) GetCSV()string{
	outputcsv := "height,txHash,Commission,Delegator_Reward\n"
	commissionSum:=0
	rewardSum:=0
	for _, b := range v {
		outputcsv += fmt.Sprintf("%s,%s,%d,%d\n",
			b.Height, b.TxHash, b.Commission, b.Reward)
			commissionSum+=b.Commission
			rewardSum+=b.Reward
	}
	outputcsv+=fmt.Sprintf("\n Sum, ,%d,%d\n",commissionSum,rewardSum)
	
	return outputcsv
}

