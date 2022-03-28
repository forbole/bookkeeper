package cosmos

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	//"regexp"
	//"strings"

	cosmostypes "github.com/forbole/bookkeeper/module/cosmos/types"
	"github.com/forbole/bookkeeper/types"
)

func GetTxs(details types.IndividualChain)error{
	for _,account := range details.FundHoldingAccount{
		txs,err:=readTxs(details.RpcEndpoint,account)
		if err!=nil{
			return err
		}
		for _,tx:=range txs.Result.Txs{
			rawlog:=strings.ReplaceAll(tx.TxResult.Log,`"{`,`{`)
			rawlog=strings.ReplaceAll(rawlog,`}"`,`}`)
			rawlog=strings.ReplaceAll(rawlog,`\`,``) 
			
			var logs []cosmostypes.RawLog
			err=json.Unmarshal([]byte(rawlog),&logs)
			if err!=nil{
				return fmt.Errorf("Error to unmarshal json object:%s",err)
			}
			fmt.Println(logs)
			bz,err:=logs[0].Events[1].Attributes[1].Value.MarshalJSON()
			if err!=nil{
				return err
			}
			fmt.Println(string(bz))
			// Seperate different message here

		}
		
		//fmt.Println(txs)
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