package flow

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	flowtypes "github.com/forbole/bookkeeper/module/flow/types"
)

func GetBalanceForEachMonth()(error){
	limit:=10
	nodeId:="21b21ad1ddb5e3002cc6a3faa55e23d70db014ee229c213f7a43769789125536"
	queryStr:=fmt.Sprintf(`{
		node_infos_from_table(limit: %d, where: {id: {_eq: "%s"}}) {
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
	request, err := http.NewRequest("POST", "https://gql.flow.forbole.com/v1/graphql", bytes.NewBuffer(jsonValue))
	client := &http.Client{Timeout: time.Second * 10}
	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode!=200{
		return fmt.Errorf("Error when getting response:%s",response.Status)
	}
	bz, _ := ioutil.ReadAll(response.Body)
	var txSearchRes flowtypes.NodeInfo
	err = json.Unmarshal(bz, &txSearchRes)
	if err != nil {
		return fmt.Errorf("Fail to marshal:%s", err)
	}
	return nil
}

