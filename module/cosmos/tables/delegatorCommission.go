package cosmos

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/forbole/bookkeeper/types"
)

func GetRewardCommission(addressBalanceEntry types.AddressBalanceEntry,denom string)(*types.AddressRewardCommission,error){
	var rewardCommission types.RewardCommissions
	for _,balanceEntry:=range addressBalanceEntry.Rows{
		instring:=strings.ReplaceAll(balanceEntry.In,denom,"")
		//fmt.Println(instring)
		in := big.NewInt(0)
		_, ok := in.SetString(instring, 10); 
		if !ok {
			return nil,fmt.Errorf("Cannot parse the int:%s",instring)
		}

		if balanceEntry.MsgType==
		"/cosmos.distribution.v1beta1.MsgWithdrawDelegatorReward" ||
		balanceEntry.MsgType==
		"withdraw_delegator_reward"{
			c:=types.NewRewardCommission(balanceEntry.TxHash,balanceEntry.Height,
				denom, big.NewInt(0),in)
			rewardCommission=append(rewardCommission,c)
		}else if balanceEntry.MsgType==
		"/cosmos.distribution.v1beta1.MsgWithdrawValidatorCommission" ||
		balanceEntry.MsgType==
		"withdraw_validator_commission"{
			c:=types.NewRewardCommission(balanceEntry.TxHash,balanceEntry.Height,
				denom,in, big.NewInt(0))
			rewardCommission=append(rewardCommission,c)
		}
	}
	addressRewardCommission:=types.NewAddressRewardCommission(addressBalanceEntry.Address,rewardCommission)
	return &addressRewardCommission,nil
}