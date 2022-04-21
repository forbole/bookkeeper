package utils

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/onflow/flow-go-sdk/client"
	"google.golang.org/grpc"
)

type FlowClient struct{
	client *client.Client
}

func NewFlowClient(accessNode string)(*FlowClient,error){
	flowClient, err := client.New(accessNode, grpc.WithInsecure())
	if err!=nil{
		return nil,err
	}
	return &(FlowClient{client:flowClient}),nil
}

func (flowClient FlowClient) GetDateByHeightMainnet(height uint64)(*time.Time,error){
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	block,err:=flowClient.client.GetBlockByHeight(ctx,height)
	if err!=nil{
		return nil,err
	}
	return &block.Timestamp,nil
}

func (flowClient FlowClient) GetDateByHeight(height uint64,lastSpork int)(*time.Time,error){
	fmt.Println(height)
	date,err:=flowClient.GetDateByHeightMainnet(height)
	for err!=nil&& strings.Contains(err.Error(),"failed to retrieve block ID for height"){
		endpoint:=fmt.Sprintf("access-001.mainnet%d.nodes.onflow.org:9000",lastSpork)
		fmt.Println(endpoint)
		newClient, err := client.New(endpoint, grpc.WithInsecure())
		if err!=nil{
			return nil,err
		}
		flowClient.client=newClient
		newDate,err:=flowClient.GetDateByHeightMainnet(height)
		lastSpork-=1
		if err==nil{
			fmt.Println(date)
			date=newDate
			break
		}
		if lastSpork==0{
			return nil,fmt.Errorf("Cannot find the block height")
		}
	}
	return date,nil
}