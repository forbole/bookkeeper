package elrond

import (
	"fmt"
	"io/ioutil"

	"github.com/forbole/bookkeeper/module/elrond/client"
	"github.com/forbole/bookkeeper/module/elrond/tables"
	"github.com/forbole/bookkeeper/types"
	"github.com/rs/zerolog/log"
)

func HandleTx(elrond types.Elrond,period types.Period,outputFolder string)([]string,error){
	log.Trace().Str("module", "elrond").Msg("HandleTx")

	var filenames []string
	client:=client.NewElrondClient(elrond.Api)
	for _,address:=range elrond.Addresses{
		entries,err:=tables.GetTxs(client,address,elrond.ValidatorContract,period.From)
		if err!=nil{
			return nil,err
		}
		csv:=entries.Rows.GetCSV()

		filename2 := fmt.Sprintf("%s/%s_%s_txs.csv", outputFolder, "elrond", address)
			err = ioutil.WriteFile(filename2, []byte(csv), 0600)
			if err != nil {
				return nil, err
			}
			filenames = append(filenames, filename2)
	}
	return nil,nil
}