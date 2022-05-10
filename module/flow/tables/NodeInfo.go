package tables

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	flowtypes "github.com/forbole/bookkeeper/module/flow/types"
)


func GetNodeInfo(nodeId string,flowjuno string)(flowtypes.NodeInfoFromTables,error){
	limit:=10
	queryStr:=fmt.Sprintf(`{
		node_infos_from_table(limit: %d, where: {id: {_eq: "%s"}}, order_by: {height: desc}) {
			id
			tokens_committed
			tokens_requested_to_unstake
			tokens_rewarded
			tokens_staked
			tokens_unstaked
			tokens_unstaking
			height
		}
	}`,limit,nodeId)
	jsonData := map[string]string{
		"query" : queryStr,
	}
	jsonValue, _ := json.Marshal(jsonData)
	request, err := http.NewRequest("POST", flowjuno, bytes.NewBuffer(jsonValue))
	client := &http.Client{Timeout: time.Second * 10}
	response, err := client.Do(request)
	if err != nil {
		return nil,err
	}
	defer response.Body.Close()

	if response.StatusCode!=200{
		return nil,fmt.Errorf("Error when getting response:%s",response.Status)
	}
	bz, _ := ioutil.ReadAll(response.Body)
	var txSearchRes flowtypes.NodeInfo
	err = json.Unmarshal(bz, &txSearchRes)
	if err != nil {
		return nil,fmt.Errorf("Fail to marshal:%s", err)
	}
	return txSearchRes.FlattenToNodeInfo(),nil
}

  func GetNodeInfoFromAddress(address string,flowjuno string,startHeight uint64)(flowtypes.NodeInfoFromTables,error){
	queryStr:=fmt.Sprintf(`{
		staker_node_id(where: {address: {_eq: %s}}) {
			staking_table {
			  node_infos_from_tables(order_by: {height: desc}, where: {height: {_gte:%d}}) {
				tokens_committed
				tokens_requested_to_unstake
				tokens_rewarded
				tokens_staked
				tokens_unstaked
				tokens_unstaking
				height
			  }
			}
		  }
	}`,address,startHeight)
	jsonData := map[string]string{
		"query" : queryStr,
	}
	jsonValue, _ := json.Marshal(jsonData)
	request, err := http.NewRequest("POST", flowjuno, bytes.NewBuffer(jsonValue))
	client := &http.Client{Timeout: time.Second * 10}
	response, err := client.Do(request)
	if err != nil {
		return nil,err
	}
	defer response.Body.Close()

	if response.StatusCode!=200{
		return nil,fmt.Errorf("Error when getting response:%s",response.Status)
	}
	bz, _ := ioutil.ReadAll(response.Body)
	var txSearchRes flowtypes.GetDataFromAddress
	err = json.Unmarshal(bz, &txSearchRes)
	if err != nil {
		return nil,fmt.Errorf("Fail to marshal:%s", err)
	}



	return txSearchRes.FlattenToNodeInfo(),nil
}


