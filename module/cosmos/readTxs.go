package cosmos

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	//"regexp"
	//"strings"

	cosmostypes "github.com/forbole/bookkeeper/module/cosmos/types"
	"github.com/forbole/bookkeeper/types"
)

// GetTxs get all the transactions from a fund raising address or self delegation address
func GetTxs(details types.IndividualChain)(types.BalanceEntries,error){
	var balanceEntries []types.BalanceEntry
	for _,address := range details.FundHoldingAccount{
		txs,err:=readTxs(details.RpcEndpoint,address)
		if err!=nil{
			return nil,err
		}
		
		for _,tx:=range txs.Result.Txs{
			
			rawlog:=strings.ReplaceAll(tx.TxResult.Log,`"{`,`{`)
			rawlog=strings.ReplaceAll(rawlog,`}"`,`}`)
			rawlog=strings.ReplaceAll(rawlog,`\n`,`,`) 
			rawlog=strings.ReplaceAll(rawlog,`\`,``) 
			
			var logs []cosmostypes.RawLog
			err=json.Unmarshal([]byte(rawlog),&logs)
			if err!=nil{
				balanceEntries=append(balanceEntries,
					types.NewBalanceEntry(tx.Height,tx.Hash,0,0,"Error reading Log for that tx"))
				continue
				/* return nil,fmt.Errorf("Error to unmarshal json object:%s\n:string:%s\n:txid:%s\n",
				err,tx.TxResult.Log,tx.Hash) */
			}
			fmt.Println(tx.Hash)
			for _,log:=range logs{
				// There will be one transaction
				fmt.Println(fmt.Sprintf("MsgIndex:%d",log.MsgIndex))
				in:=0
				out:=0
				msgType:=""

				// Read event for that log
				for _,event:=range log.Events{
					//Catagorise each event and put it in a table
					attribute:=ConvertAttributeToMap(event.Attributes)
					//fmt.Println(fmt.Sprintf("type:%s",event.Type))
					// check if we are the receiver (write on + side)
					if event.Type=="coin_received"{
						bz,err:=attribute["receiver"].MarshalJSON()
						if err!=nil{
							return nil,err
						}
						receiver:=string(bz)
						receiver=strings.ReplaceAll(receiver,"\"","")
						
						if receiver!=address{
							continue
						}  
						bzamount,err:=attribute["amount"].MarshalJSON()
						if err!=nil{
							return nil,err
						}
						amount:=strings.ReplaceAll(string(bzamount),"\"","")
						amount=strings.ReplaceAll(amount,details.Denom,"")
						a,err:=strconv.Atoi(amount)
						if err!=nil{
							return nil,err
						}
						in=a
						
						fmt.Println(fmt.Sprintf("Receive amount:%d\nReceiver:%s",in,receiver))
					}
					if event.Type=="coin_spent"{
						bz,err:=attribute["spender"].MarshalJSON()
						if err!=nil{
							return nil,err
						}

						spender:=string(bz)
						spender=strings.ReplaceAll(spender,"\"","")
						if string(spender)!=address{
							continue
						}
						bzamount,err:=attribute["amount"].MarshalJSON()
						if err!=nil{
							return nil,err
						}
						amount:=string(bzamount)
						amount=strings.ReplaceAll(amount,"\"","")
						amount=strings.ReplaceAll(amount,details.Denom,"")
						b,err:=strconv.Atoi(amount)
						if err!=nil{
							return nil,err
						}
						out=b

						fmt.Println(fmt.Sprintf("Spent amount:%d\nspender:%s",out,spender))
					}
					if event.Type=="message"{
						bzaction,err:=attribute["action"].MarshalJSON()
						if err!=nil{
							return nil,err
						}
						fmt.Println(fmt.Sprintf("action:%s",string(bzaction)))
						msgType=string(bzaction)

					}
				}

				balanceEntries=append(balanceEntries,
					types.NewBalanceEntry(tx.Height,tx.Hash,in,out,msgType))

			}
		}
	}
	return balanceEntries,nil
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