package flow

import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/forbole/bookkeeper/module/flow/tables"
	"github.com/forbole/bookkeeper/module/flow/utils"

	"github.com/forbole/bookkeeper/types"
)

func HandleNodeInfos(flow types.Flow, vsCurrency string, period types.Period) ([]string, error) {
	if len(flow.Addresses) == 0 {
		return nil, nil
	}
	var filenames []string
	flowClient, err := utils.NewFlowClient(flow.FlowEndpoint)
	if err != nil {
		return nil, err
	}

	startDate := time.Unix(period.From, 0)
	startHeight, err := flowClient.GetHeightByDate(startDate, flow.LastSpork)
	if err != nil {
		return nil, err
	}

	for _, id := range flow.Addresses {
		nodeInfo, err := tables.GetNodeInfoFromAddress(id, flow.FlowJuno, startHeight)
		if err != nil {
			return nil, err
		}

		outputcsv, err := nodeInfo.GetCSV(flow.Exponent, "flow", vsCurrency, *flowClient, flow.LastSpork)
		if err != nil {
			return nil, err
		}
		fmt.Println(outputcsv)
		filename := fmt.Sprintf("%s_nodeInfo.csv", id)
		err = ioutil.WriteFile(filename, []byte(outputcsv), 0777)
		if err != nil {
			return nil, err
		}
		filenames = append(filenames, filename)
	}
	return filenames, nil

}
