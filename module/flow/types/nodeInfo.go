package types

import (
	"fmt"
	"math"
	"net/http"
	"time"

	"github.com/forbole/bookkeeper/module/flow/utils"
	coingecko "github.com/superoo7/go-gecko/v3"

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

func (n NodeInfo) GetCSV(exp int,coinId string,vsCurrency string,flowclient utils.FlowClient)(string,error){
	outputcsv := "Date,TokensCommitted,TokensRequestedToUnstake,TokensRewarded,TokensStaked,TokensUnstaked,TokensUnstaking, ,TokensCommitted_converted,TokensRequestedToUnstake_converted,TokensRewarded_converted,TokensStaked_converted,TokensUnstaked_converted,TokensUnstaking_converted\n"
	exponent := math.Pow(10, float64(-1*exp))

	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}
	
	CG := coingecko.NewClient(httpClient)

	singlePrice,err:=CG.CoinsMarket(vsCurrency, []string{coinId}, "", 0, 0, false, nil)
	if err!=nil{
		return "",err
	}
	fmt.Println((*singlePrice)[0].CurrentPrice)
	coinprice:=float64((*singlePrice)[0].CurrentPrice)
	fmt.Println(coinprice)


	for _, b := range n.Data.NodeInfosFromTable {
		date,err:=flowclient.GetDateByHeight(uint64(b.Height))
		if err!=nil{
			return "",err
		}

		commited:=float64(b.TokensCommitted)*exponent
		requestToUnstake:= float64(b.TokensRequestedToUnstake)*exponent
		rewarded:=float64(b.TokensRewarded)*exponent
		staked:=float64(b.TokensStaked)*exponent
		unstaked:=float64(b.TokensUnstaked)*exponent
		unstaking:=float64(b.TokensUnstaking)*exponent
		outputcsv += fmt.Sprintf("%s,%f,%f,%f,%f,%f,%f, ,%f,%f,%f,%f,%f,%f\n",
			date, 
			commited,requestToUnstake,
			rewarded,staked,
			unstaked,unstaking,
			commited*coinprice,requestToUnstake*coinprice,
			rewarded*coinprice,staked*coinprice,
			unstaked*coinprice,unstaking*coinprice,)

	}	
	return outputcsv,nil
}