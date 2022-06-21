package tables

import (
	"fmt"

	elrondClient "github.com/forbole/bookkeeper/module/elrond/client"
	"github.com/forbole/bookkeeper/types/tabletypes"
	"github.com/rs/zerolog/log"
)

// GetTxs 
func GetTxs(client elrondClient.ElrondClient,address,contract string,from int64)(*tabletypes.AddressBalanceEntry,error){
	log.Trace().Str("module", "elrond").Msg("GetTxs")

	txs,err:=client.GetTxs(address,"")
	if err!=nil{
		return nil,err
	}
	fmt.Println(len(txs))

	var balanceEntries []tabletypes.BalanceEntry
	for _,tx:=range txs{
		if tx.Status=="fail"{
			continue
		}
		txResult,err:=client.GetTxResult(tx.TxHash)
		if err!=nil{
			return nil,err
		}

		for _,result:=range txResult.Results{
			balanceEntries=append(balanceEntries,
				tabletypes.NewBalanceEntry(tx.Round,tx.TxHash,result.Value,"0",tx.Function))
		}
	}

	accountBalanceSheet:=tabletypes.NewAccountBalanceSheet(address,balanceEntries)
	return &accountBalanceSheet,nil
}