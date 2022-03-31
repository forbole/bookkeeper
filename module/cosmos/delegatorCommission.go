package cosmos

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/forbole/bookkeeper/types"
)

func GetRewardCommission(addressBalanceEntry types.AddressBalanceEntry,denom string)(*types.AddressRewardCommission,error){
	var rewardCommission types.RewardCommissions
	for _,balanceEntry:=range addressBalanceEntry.BalanceEntry{
		instring:=strings.ReplaceAll(balanceEntry.In,denom,"")
		fmt.Println(instring)
		in,err:=strconv.Atoi(instring)
		fmt.Println(balanceEntry.MsgType)
		if err!=nil{
			return nil,err
		}

		if balanceEntry.MsgType==
		"/cosmos.distribution.v1beta1.MsgWithdrawDelegatorReward" ||
		balanceEntry.MsgType==
		"withdraw_delegator_reward"{
			c:=types.NewRewardCommission(balanceEntry.TxHash,balanceEntry.Height,
				denom,0,in)
			rewardCommission=append(rewardCommission,c)
		}else if balanceEntry.MsgType==
		"/cosmos.distribution.v1beta1.MsgWithdrawValidatorCommission" ||
		balanceEntry.MsgType==
		"withdraw_validator_commission"{
			c:=types.NewRewardCommission(balanceEntry.TxHash,balanceEntry.Height,
				denom,in,0)
			rewardCommission=append(rewardCommission,c)
		}
	}
	addressRewardCommission:=types.NewAddressRewardCommission(addressBalanceEntry.Address,rewardCommission)
	fmt.Println(addressRewardCommission)
	return &addressRewardCommission,nil
}