package tables

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/forbole/bookkeeper/coinApi"
	"github.com/forbole/bookkeeper/module/flow/db"
	flowutils "github.com/forbole/bookkeeper/module/flow/utils"
	"github.com/forbole/bookkeeper/types/tabletypes"
)

// GetRewardCommission get the reward from the db and turn it into rewardcommisson struct
func GetRewardCommission(payer string,db *db.FlowDb,flowClient *flowutils.FlowClient,vsCurrency string)(*tabletypes.AddressDateRewardPrice,error){
	rewardRaw,err:=db.GetWithdrawReward(payer)
	if err!=nil{
		return nil,err
	}

	rewardRow:=make([]tabletypes.DateRewardPriceRow,len(rewardRaw))
	for i,r:=range rewardRaw{
		date,err:=flowClient.GetDateByHeight(uint64(r.Height))
		if err!=nil{
			return nil,err
		}

		price,err:=coinApi.GetCryptoPriceFromDate(*date,"flow",vsCurrency)
		if err!=nil{
			return nil,err
		}
		
		rewardIndex:=strings.Index(r.Value,"amount: ")
		rewardStr:=r.Value[rewardIndex:len(r.Value)-2]
		fmt.Println(rewardStr)
		reward,ok:=new(big.Float).SetString(rewardStr)
		if !ok{
			return nil,fmt.Errorf("Cannot convert to big.Float from string")
		}
		rewardPrice := new(big.Float).Mul(reward, price)

//A.8624b52f9ddcd04a.FlowIDTableStaking.RewardTokensWithdrawn(nodeID: "237a7a04ecf88b7c21001589ecc277190a6f7cd6e56a296a203552ade6db0927", amount: 1526.81000000)
		rewardRow[i]=tabletypes.NewDateRewardPriceRow(*date,reward,new(big.Float).SetInt64(0),"flow",rewardPrice,new(big.Float).SetInt64(0))
	}

	return &tabletypes.AddressDateRewardPrice{Address: payer,Rows: rewardRow},nil
}
