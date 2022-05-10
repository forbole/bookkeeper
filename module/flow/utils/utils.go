package utils

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/onflow/flow-go-sdk/client"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type FlowClient struct {
	client *client.Client
}

func NewFlowClient(accessNode string) (*FlowClient, error) {
	flowClient, err := client.New(accessNode, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return &(FlowClient{client: flowClient}), nil
}

func (flowClient FlowClient) GetDateByHeightMainnet(height uint64) (*time.Time, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	block, err := flowClient.client.GetBlockByHeight(ctx, height)
	if err != nil {
		return nil, err
	}
	return &block.Timestamp, nil
}

// GetDateByHeight get the time for the specific height
func (flowClient FlowClient) GetDateByHeight(height uint64, lastSpork int) (*time.Time, error) {
	//fmt.Println(height)
	date, err := flowClient.GetDateByHeightMainnet(height)
	for err != nil && strings.Contains(err.Error(), "failed to retrieve block ID for height") {
		endpoint := fmt.Sprintf("access-001.mainnet%d.nodes.onflow.org:9000", lastSpork)
		//fmt.Println(endpoint)
		newClient, err := client.New(endpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			return nil, err
		}
		flowClient.client = newClient
		newDate, err := flowClient.GetDateByHeightMainnet(height)
		lastSpork--
		if err == nil {
			//fmt.Println(date)
			date = newDate
			break
		}
		if lastSpork == 0 {
			return nil, fmt.Errorf("cannot find the block height")
		}
	}
	return date, nil
}

func (flowClient FlowClient) GetHeightByDate(t time.Time, lastSpork int) (uint64, error) {
	// GetHeightByDate get height for the cloest time stamp within 10 seconds, log2 complexity
	// If the height do not exist, will return the Lowest Height
	// if it is before network start, return first block height
	if t.Before(time.Date(2020, 10, 13, 0, 0, 0, 0, time.UTC)) {
		return 7601063, nil
	}

	// get latest height
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	block, err := flowClient.client.GetLatestBlock(ctx, true)
	if err != nil {
		return 0, err
	}

	latestHeight := block.Height

	left := latestHeight
	newT := block.Timestamp
	right := uint64(7601063)
	middle := (left + right) / 2
	if err != nil {
		return 0, err
	}
	//fmt.Println(t)

	for !((t.After(newT) && t.Sub(newT) < (time.Hour*23)) ||
		(t.Before(newT) && newT.Sub(t) < (time.Hour*23))) {
		middle = (left + right) / 2

		if middle == 1 {
			// 1 is the earliest block time
			return 1, nil
		}

		T, err := flowClient.GetDateByHeight(middle, 16)
		if err != nil {
			return 0, err
		}

		newT = *T
		if newT.After(t) {
			left = middle + 1
		} else if newT.Before(t) {
			right = middle - 1
		}
	}
	return middle, nil

}
