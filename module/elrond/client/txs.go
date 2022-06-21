package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	elrondtypes "github.com/forbole/bookkeeper/module/elrond/types"
	"github.com/rs/zerolog/log"
)

type ElrondClient struct{
	api string
}

func NewElrondClient(api string)ElrondClient{
	return ElrondClient{
		api: api,
	}
}

// GetSelfRedelegate get transaction for the speific contract
func (client *ElrondClient)GetSelfRedelegate(address, contract string,from int64)([]elrondtypes.Tx,error){
	log.Trace().Str("module", "elrond").Msg("GetSelfRedelegate")

	parameter:=fmt.Sprintf("&sender=%s&receiver=%s",
		address,contract)
		
	return client.GetTxs(address,parameter)
	
}

// GetTxs get transaction for address and given parameter string
func (client *ElrondClient)GetTxs(address string,parameters string)([]elrondtypes.Tx,error){
	log.Trace().Str("module", "elrond").Msg("GetTxs")

	var elrondTxs []elrondtypes.Tx
	size:=50
	for i:=0;;i++{

		query := fmt.Sprintf("accounts/%s/transactions?size=%d&from=%d%s",
		 address,size, i*size,parameters)

		bz,err:=client.ping(query)

		var txs []elrondtypes.Tx
		err = json.Unmarshal(bz, &txs)
		if err != nil {
			return nil, fmt.Errorf("Fail to marshal:%s", err)
		}

		elrondTxs=append(elrondTxs, txs...)
		if len(txs)<size{
			break
		}

	}
	
	return elrondTxs,nil
}

func (client *ElrondClient)GetTxResult(txHash string)(*elrondtypes.TxResult,error){
	query := fmt.Sprintf("transactions/%s",txHash)

	bz,err:=client.ping(query)

	var tx elrondtypes.TxResult
	err = json.Unmarshal(bz, &tx)
	if err != nil {
		return nil, fmt.Errorf("Fail to marshal:%s", err)
	}
	return &tx,nil
}

func (client *ElrondClient)ping(query string)([]byte,error){
	
	q := fmt.Sprintf("%s/%s",client.api,query)

	fmt.Println(q)
	var bz []byte
	resp, err := http.Get(q)
	if err != nil {
		return nil, fmt.Errorf("Fail to get tx from rpc:%s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Fail to get tx from rpc:Status :%s", resp.Status)
	}

	bz, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return bz,nil
}
