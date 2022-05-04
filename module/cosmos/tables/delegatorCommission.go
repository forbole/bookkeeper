package tables

import (
	"fmt"
	"math/big"
	"strings"

	tabletypes "github.com/forbole/bookkeeper/types/tabletypes"
)

// GetRewardCommission filter the rewardCommission table for balance entry
func GetRewardCommission(addressBalanceEntry tabletypes.AddressBalanceEntry)(*tabletypes.AddressRewardCommission,error){
	var rewardCommission tabletypes.RewardCommissions

	for _,balanceEntry:=range addressBalanceEntry.Rows{
		denomArray:=strings.Split(balanceEntry.In,",")
		for _,denomString:=range denomArray{
			//fmt.Println(denomString)
			in,denom,err:=splitValueDenom(denomString)
			if err!=nil{
				return nil,err
			}

			if balanceEntry.MsgType==
				"/cosmos.distribution.v1beta1.MsgWithdrawDelegatorReward" ||
			balanceEntry.MsgType==
				"withdraw_delegator_reward"{

				c:=tabletypes.NewRewardCommission(balanceEntry.TxHash,balanceEntry.Height,
					denom, big.NewInt(0),in)
				rewardCommission=append(rewardCommission,c)

			}else if balanceEntry.MsgType==
				"/cosmos.distribution.v1beta1.MsgWithdrawValidatorCommission" ||
			balanceEntry.MsgType==
				"withdraw_validator_commission"{

				c:=tabletypes.NewRewardCommission(balanceEntry.TxHash,balanceEntry.Height,
					denom,in, big.NewInt(0))
				rewardCommission=append(rewardCommission,c)

			}
		
		
		}
		
	}
	addressRewardCommission:=tabletypes.NewAddressRewardCommission(addressBalanceEntry.Address,rewardCommission)
	return &addressRewardCommission,nil
}

// splitValueDenom split string contains value and denom
// eg."1523uatom" -> 1523, uatom, nil
func splitValueDenom(value string)(*big.Int, string, error){
	// split
		var l, n []rune
		for _, r := range value {
			switch {
			case r >= 'A' && r <= 'Z':
				l = append(l, r)
			case r >= 'a' && r <= 'z':
				l = append(l, r)
			case r >= '0' && r <= '9':
				n = append(n, r)
			}
		}

		in := big.NewInt(0)
		_, ok := in.SetString(string(n), 10); 
		if !ok {
			return nil,"",fmt.Errorf("Cannot parse the int:%s",string(n))
		}
		return in,string(l),nil
}