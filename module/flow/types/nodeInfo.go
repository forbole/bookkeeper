package types

import (
	"fmt"
	"math"

	"github.com/forbole/bookkeeper/module/flow/utils"
)

type NodeInfo struct {
	Data struct {
		NodeInfosFromTable []struct {
			ID                       string `json:"id"`
			TokensCommitted          int64  `json:"tokens_committed"`
			TokensRequestedToUnstake int64    `json:"tokens_requested_to_unstake"`
			TokensRewarded           int64    `json:"tokens_rewarded"`
			TokensStaked             int64  `json:"tokens_staked"`
			TokensUnstaked           int64    `json:"tokens_unstaked"`
			TokensUnstaking          int64    `json:"tokens_unstaking"`
			Height                   int64    `json:"height"`
		} `json:"node_infos_from_table"`
	} `json:"data"`
}

func (n NodeInfo) GetCSV(exp int,denom string,flowclient utils.FlowClient)(string,error){
	outputcsv := "Date,TokensCommitted,TokensRequestedToUnstake,TokensRewarded,TokensStaked,TokensUnstaked,TokensUnstaking\n"
	commissionSum:=0
	rewardSum:=0
	exponent := math.Pow(10, float64(-1*exp))


	for _, b := range n.Data.NodeInfosFromTable {
		date,err:=flowclient.GetDateByHeight(uint64(b.Height))
		if err!=nil{
			return "",err
		}
		
		outputcsv += fmt.Sprintf("%s,%f,%f,%f,%f,%f,%f\n",
			date, 
			float64(b.TokensCommitted)*exponent, float64(b.TokensRequestedToUnstake)*exponent,
			float64(b.TokensRewarded)*exponent,float64(b.TokensStaked)*exponent,
			float64(b.TokensUnstaked)*exponent,float64(b.TokensUnstaking)*exponent)
	}

	outputcsv+=fmt.Sprintf("\n Sum, ,%f,%f\n",float64(commissionSum)*exponent,float64(rewardSum)*exponent)
	
	return outputcsv,nil
}