package client
import (
	tmrpc "github.com/tendermint/tendermint/rpc/client/http"
	"google.golang.org/grpc"
	"context"

)
type CosmosClient struct{
	chainId string
	
}

func CreateCosmosGrpcClient(grpcAddress string,rpcAddress string){
	grpcConn, err := grpc.Dial(
		grpcAddress,
		grpc.WithInsecure(),
	)
	if err != nil {
		panic(err)
	}
	defer grpcConn.Close()

	chainID := getChainID(rpcAddress)
}

func getChainID(rpc string) string {
	client, err := tmrpc.New(rpc, "/websocket")
	if err != nil {
		panic(err)
	}

	status, err := client.Status(context.Background())
	if err != nil {
		panic(err)
	}

	return status.NodeInfo.Network
}

func GetAccount()