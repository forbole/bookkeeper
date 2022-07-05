package solana

import (
	"fmt"

	"github.com/forbole/bookkeeper/module/solana/client"
	"github.com/forbole/bookkeeper/module/solana/tables"
	"github.com/forbole/bookkeeper/types"
)

func HandleReward(solana types.Solana,period types.Period, vsCurrency string)([]string,error){
	solClient:=client.NewSolanaBeachClient(solana.SolanaBeachApi)
	addressRewardPriceTables,err:=tables.GetStakeRewardForPubKey(solana,period.From,vsCurrency,solClient)
	if err!=nil{
		return nil,err
	}
	for _,table:=range addressRewardPriceTables{
		fmt.Println(table.GetCSV())
	}
	return nil,nil
}