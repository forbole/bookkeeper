package cosmos

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	cosmostypes "github.com/forbole/bookkeeper/module/cosmos/types"
	"github.com/forbole/bookkeeper/types"
)

func GetTxs(details types.IndividualChain)error{
	for _,account := range details.FundHoldingAccount{
		txs,err:=readTxs(details.RpcEndpoint,account)
		if err!=nil{
			return err
		}
		txs.
		fmt.Println(txs)
	}
	return nil
}

func readTxs(api string, address string)(*cosmostypes.TxSearchRespond,error){
	limit:=2
	page:=1
	query:=fmt.Sprintf(`%s/tx_search?query="message.sender='%s'"&prove=true&page=%d&per_page=%d&order_by="desc"`,
	api,address,page,limit)
	fmt.Println(query)
	resp,err:=http.Get(query)
	if err!=nil{
		return nil,fmt.Errorf("Fail to get tx from rpc:%s",err)
	}
	if resp.StatusCode!=200 {
		return nil,fmt.Errorf("Fail to get tx from rpc:Status code:%d",resp.StatusCode)
	}

	defer resp.Body.Close()

	bz, err := io.ReadAll(resp.Body)
	

	var txSearchRes cosmostypes.TxSearchRespond
	err=json.Unmarshal(bz,&txSearchRes)
	if err!=nil{
		return nil,fmt.Errorf("Fail to marshal:%s",err)
	}

	


	
	return &txSearchRes,nil
}