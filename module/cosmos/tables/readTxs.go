package cosmos

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"strconv"
	"strings"

	//"regexp"
	//"strings"

	cosmostypes "github.com/forbole/bookkeeper/module/cosmos/types"
	"github.com/forbole/bookkeeper/types"
)

const (
	//height at 2021-01-01 12:01:14
	height = 4635000
)

// GetTxs get all the transactions from a fund raising address or self delegation address
func GetTxs(details types.IndividualChain) ([]types.AddressBalanceEntry, error) {
	var accountbalanceEntries []types.AddressBalanceEntry

	for _, address := range details.FundHoldingAccount {
		var balanceEntries []types.BalanceEntry
		res, err := readTxs(details.RpcEndpoint, address, height)
		if err != nil {
			return nil, err
		}
		if res==nil{
			return nil,nil
		}
		for _, txs := range res {
			for _, tx := range txs.Result.Txs {

				rawlog := strings.ReplaceAll(tx.TxResult.Log, `"{`, `{`)
				rawlog = strings.ReplaceAll(rawlog, `}"`, `}`)
				rawlog = strings.ReplaceAll(rawlog, `\n`, `,`)
				rawlog = strings.ReplaceAll(rawlog, `\`, ``)

				var logs []cosmostypes.RawLog
				err = json.Unmarshal([]byte(rawlog), &logs)
				height,err:=strconv.Atoi(tx.Height)
				if err!=nil{
					return nil,err
				}
				if err != nil {
					balanceEntries = append(balanceEntries,
						types.NewBalanceEntry(height, tx.Hash, "0", "0", "Error reading Log for that tx"))
					continue
					/* return nil,fmt.Errorf("Error to unmarshal json object:%s\n:string:%s\n:txid:%s\n",
					err,tx.TxResult.Log,tx.Hash) */
				}
				fmt.Println(tx.Hash)

				balanceEntry,err:=readlogs(logs,address,tx.Hash,height)
				if err!=nil{
					return nil,err
				}

				balanceEntries = append(balanceEntries,
					balanceEntry...)
				}
			}
			accountbalanceEntries = append(accountbalanceEntries,
				types.NewAccountBalanceSheet(address, balanceEntries))
		}
	
	return accountbalanceEntries, nil
}

func readlogs(logs []cosmostypes.RawLog,address,hash string,height int)([]types.BalanceEntry,error){
	var balanceEntries []types.BalanceEntry

	for _, log := range logs {
		// There will be one transaction
		fmt.Println(fmt.Sprintf("MsgIndex:%d", log.MsgIndex))
		in := "0"
		out := "0"
		msgType := ""

		// Read event for that log
		for _, event := range log.Events {
			//Catagorise each event and put it in a table
			attribute := ConvertAttributeToMap(event.Attributes)
			//fmt.Println(fmt.Sprintf("type:%s",event.Type))
			// check if we are the receiver (write on + side)
			if event.Type == "transfer" {
				// get the amount for the transafer
				bzamount, err := attribute["amount"].MarshalJSON()
				if err != nil {
					return nil, err
				}
				amount := strings.ReplaceAll(string(bzamount), "\"", "")
				//amount=strings.ReplaceAll(amount,details.Denom,"")
				//a,err:=strconv.Atoi(amount)
				if err != nil {
					return nil, err
				}
				in = amount

				// check if recipient is transfer
				bz, err := attribute["recipient"].MarshalJSON()
				if err != nil {
					return nil, err
				}
				receiver := string(bz)
				receiver = strings.ReplaceAll(receiver, "\"", "")

				if receiver == address {
					in = amount
				} else {

					bz, err = attribute["spender"].MarshalJSON()
					if err != nil {
						return nil, err
					}

					spender := string(bz)
					spender = strings.ReplaceAll(spender, "\"", "")
					if string(spender) == address {
						out = amount
					}
				}
				if in != "0" {
					fmt.Println(fmt.Sprintf("Received amount:%s\nReceiver:%s", in, receiver))
				} else if out != "0" {
					fmt.Println(fmt.Sprintf("Spent amount:%s\nSpender:%s", out, receiver))
				}
			}
			if event.Type == "message" {
				bzaction, err := attribute["action"].MarshalJSON()
				if err != nil {
					return nil, err
				}
				msgType=strings.ReplaceAll(string(bzaction),"\"","")
				fmt.Println(fmt.Sprintf("action:%s", msgType))
			}
		}
		balanceEntries = append(balanceEntries,
			types.NewBalanceEntry(height, hash, in, out, msgType))
	}

	
	return balanceEntries,nil
}

// ConvertAttributeToMap turn attribute into a map so that it is easy to find attributes
func ConvertAttributeToMap(array []cosmostypes.Attributes) map[string]json.RawMessage {
	resultingMap := map[string]json.RawMessage{}
	for _, m := range array {
		resultingMap[m.Key] = m.Value
	}
	return resultingMap
}

// readtxs read the height and read to the page that meet the target height
func readTxs(api string, address string, targetHeight int) ([]*cosmostypes.TxSearchRespond, error) {
	var res []*cosmostypes.TxSearchRespond
	limit := 30
	lastHeight := math.MaxInt
	totalCount := math.MaxInt
	for page := 1; (lastHeight > targetHeight) && ((limit*page)<totalCount); page++ {
		query := fmt.Sprintf(`%s/tx_search?query="message.sender='%s'"&prove=true&page=%d&per_page=%d&order_by="desc"`,
			api, address, page, limit)
		fmt.Println(query)
		resp, err := http.Get(query)
		if err != nil {
			return nil, fmt.Errorf("Fail to get tx from rpc:%s", err)
		}
		if resp.StatusCode != 200 {
			return nil, fmt.Errorf("Fail to get tx from rpc:Status :%s", resp.Status)
		}

		defer resp.Body.Close()

		bz, err := io.ReadAll(resp.Body)

		var txSearchRes cosmostypes.TxSearchRespond
		err = json.Unmarshal(bz, &txSearchRes)
		if err != nil {
			return nil, fmt.Errorf("Fail to marshal:%s", err)
		}
		if txSearchRes.Result.TotalCount=="0"{
			return nil,nil
		}
		lastHeight, err = strconv.Atoi(
			txSearchRes.Result.Txs[len(txSearchRes.Result.Txs)-1].Height,
		)
		if err != nil {
			return nil, err
		}
		totalCount,err=strconv.Atoi(txSearchRes.Result.TotalCount)
		if err!=nil{
			return nil,err
		}
		res = append(res, &txSearchRes)
	}

	return res, nil
}
