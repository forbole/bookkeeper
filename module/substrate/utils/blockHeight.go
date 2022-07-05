package utils

import (
	"time"

	"github.com/forbole/bookkeeper/module/substrate/client"
	substratetypes "github.com/forbole/bookkeeper/module/substrate/types"
)

func GetTimeByBlockNum(blockNum int, api *client.SubscanClient) (*time.Time, error) {
	requestUrl := "/api/scan/block"
	type Payload struct {
		BlockNum int `json:"block_num"`
	}

	payload := Payload{
		BlockNum: blockNum,
	}

	var block substratetypes.Block
	err := api.CallApi(requestUrl, payload, &block)
	if err != nil {
		return nil, err
	}

	timestamp := time.Unix(int64(block.Data.BlockTimestamp), 0)

	return &timestamp, nil
}
