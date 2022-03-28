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
			rawlog=strings.ReplaceAll(rawlog,`\n`,`,`) 
			rawlog=strings.ReplaceAll(rawlog,`\`,``) 
			
			var logs []cosmostypes.RawLog
			err=json.Unmarshal([]byte(rawlog),&logs)
			if err!=nil{
				return fmt.Errorf("Error to unmarshal json object:%s\n:string:%s\n:txid:%s\n",
				err,tx.TxResult.Log,tx.Hash)
			}

			if err!=nil{
				return err
			}
			fmt.Println(tx.Hash)
			for _,log:=range logs{
				for _,event:=range log.Events{
					//Catagorise each event and put it in a table
					attribute:=ConvertAttributeToMap(event.Attributes)
					bz,err:=attribute["action"].MarshalJSON()
					if err!=nil{
						return err
					}
					fmt.Println(string(bz))
						switch string(bz) {
						case "delegate":
							
						}
				}
			}
		}
	}
	return nil
}

// ConvertAttributeToMap turn attribute into a map so that it is easy to find attributes
func ConvertAttributeToMap(array []cosmostypes.Attributes)map[string]json.RawMessage{
    resultingMap := map[string]json.RawMessage{}
    for _, m := range array {
            resultingMap[m.Key] = m.Value
        }
	return resultingMap
}

func readTxs(api string, address string)(*cosmostypes.TxSearchRespond,error){
	limit:=30
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