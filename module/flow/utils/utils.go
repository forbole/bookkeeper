package utils

import (
	"context"
	"time"

	"github.com/onflow/flow-go-sdk"
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

func (flowClient FlowClient) GetDateByHeight(height uint64)(*time.Time,error){
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	block,err:=flowClient.client.GetBlockByHeight(ctx,height,nil)
	if err!=nil{
		return nil,err
	}
	return &block.Timestamp,nil
}