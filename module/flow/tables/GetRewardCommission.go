package tables

import (
	"github.com/forbole/bookkeeper/module/flow/db"
	"github.com/forbole/bookkeeper/types/tabletypes"


)

// GetRewardCommission get the reward from the db and turn it into rewardcommisson struct
func GetRewardCommission(payer string,db *db.FlowDb)(*tabletypes.AddressRewardCommission,error){
	rewardRaw,err:=db.GetWithdrawReward(payer)
	if err!=nil{
		return nil,err
	}

	rewardCommission:=make(tabletypes.RewardCommission)
	for 

}