package utils

import (
	"time"

	"github.com/forbole/bookkeeper/module/subtrate/client"
	subtratetypes "github.com/forbole/bookkeeper/module/subtrate/types"
)

func GetTimeByBlockNum(blockNum int, api *client.SubscanClient) (*time.Time, error) {
	requestUrl := "/api/scan/block"
	type Payload struct {
		BlockNum int `json:"block_num"`
	}

	payload := Payload{
		BlockNum: blockNum,
	}

	var block subtratetypes.Block
	err := api.CallApi(requestUrl, payload, &block)
	if err != nil {
		return nil, err
	}

	timestamp := time.Unix(int64(block.Data.BlockTimestamp), 0)

	return &timestamp, nil
}
