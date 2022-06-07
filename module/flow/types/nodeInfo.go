package types

import (
	"fmt"
	"math"
	"net/http"
	"time"

	"github.com/forbole/bookkeeper/module/flow/utils"
	coingecko "github.com/superoo7/go-gecko/v3"
)

type NodeInterface interface {
	FlattenToNodeInfo() []NodeInfoFromTable
}

type NodeInfo struct {
	Data struct {
		NodeInfosFromTable []NodeInfoFromTable `json:"node_infos_from_table"`
	} `json:"data"`
}

type NodeInfoFromTable struct {
	ID                       string `json:"id"`
	TokensCommitted          int64  `json:"tokens_committed"`
	TokensRequestedToUnstake int64  `json:"tokens_requested_to_unstake"`
	TokensRewarded           int64  `json:"tokens_rewarded"`
	TokensStaked             int64  `json:"tokens_staked"`
	TokensUnstaked           int64  `json:"tokens_unstaked"`
	TokensUnstaking          int64  `json:"tokens_unstaking"`
	Height                   int64  `json:"height"`
}

type GetDataFromAddress struct {
	Data struct {
		StakerNodeID []struct {
			StakingTable struct {
				NodeInfosFromTables []NodeInfoFromTable `json:"node_infos_from_tables"`
			} `json:"staking_table"`
		} `json:"staker_node_id"`
	} `json:"data"`
}

func (n GetDataFromAddress) FlattenToNodeInfo() []NodeInfoFromTable {
	return n.Data.StakerNodeID[0].StakingTable.NodeInfosFromTables
}

func (n NodeInfo) FlattenToNodeInfo() []NodeInfoFromTable {
	return n.Data.NodeInfosFromTable
}

type NodeInfoFromTables []NodeInfoFromTable

func (n NodeInfoFromTables) GetCSV(exp int, coinId string, vsCurrency string, flowclient utils.FlowClient, lastspork int) (string, error) {
	outputcsv := "Date,TokensCommitted,TokensRequestedToUnstake,TokensRewarded,TokensStaked,TokensUnstaked,TokensUnstaking, ,TokensCommitted_value,TokensRequestedToUnstake_value,TokensRewarded_value,TokensStaked_value,TokensUnstaked_value,TokensUnstaking_value\n"
	exponent := math.Pow(10, float64(-1*exp))

	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}

	cg := coingecko.NewClient(httpClient)

	singlePrice, err := cg.CoinsMarket(vsCurrency, []string{coinId}, "", 0, 0, false, nil)
	if err != nil {
		return "", err
	}
	//fmt.Println((*singlePrice)[0].CurrentPrice)
	coinprice := (*singlePrice)[0].CurrentPrice
	//fmt.Println(coinprice)

	for _, b := range n {
		//fmt.Println(flowclient)
		date, err := (flowclient).GetDateByHeight(uint64(b.Height))
		if err != nil {
			return "", err
		}

		committed := float64(b.TokensCommitted) * exponent
		requestToUnstake := float64(b.TokensRequestedToUnstake) * exponent
		rewarded := float64(b.TokensRewarded) * exponent
		staked := float64(b.TokensStaked) * exponent
		unstaked := float64(b.TokensUnstaked) * exponent
		unstaking := float64(b.TokensUnstaking) * exponent
		outputcsv += fmt.Sprintf("%s,%f,%f,%f,%f,%f,%f, ,%f,%f,%f,%f,%f,%f\n",
			date,
			committed, requestToUnstake,
			rewarded, staked,
			unstaked, unstaking,
			committed*coinprice, requestToUnstake*coinprice,
			rewarded*coinprice, staked*coinprice,
			unstaked*coinprice, unstaking*coinprice)

	}
	return outputcsv, nil
}
