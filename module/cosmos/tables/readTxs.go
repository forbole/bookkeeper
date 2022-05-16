package tables

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"

	//"regexp"
	//"strings"

	"github.com/rs/zerolog/log"
	cosmostypes "github.com/forbole/bookkeeper/module/cosmos/types"
	"github.com/forbole/bookkeeper/module/cosmos/utils"
	types "github.com/forbole/bookkeeper/types"
	tabletypes "github.com/forbole/bookkeeper/types/tabletypes"
)

// GetTxs get all the transactions from a fund raising address or self delegation address
func GetTxs(details types.IndividualChain, from int64) ([]tabletypes.AddressBalanceEntry, error) {
	log.Trace().Str("module", "cosmos").Msg("get txs")

	var accountbalanceEntries []tabletypes.AddressBalanceEntry
	targetHeight, err := utils.GetHeightByDate(time.Unix(from, 0), details.LcdEndpoint)
	if err != nil {
		return nil, err
	}

	for _, address := range details.FundHoldingAccount {
		accountBalanceSheet,err:=GetTxsForAnAddress(address,details.RpcEndpoint,targetHeight)
		if err!=nil{
			return nil,err
		}
		accountbalanceEntries = append(accountbalanceEntries,*accountBalanceSheet)
	}

	return accountbalanceEntries, nil
}

// GetTxsForAnAddress get txs for a single address from from now to the target height
func GetTxsForAnAddress(address string, rpcEndpoint string, targetHeight int)(*tabletypes.AddressBalanceEntry,error){
	var balanceEntries []tabletypes.BalanceEntry
	res, err := readTxs(rpcEndpoint, address, targetHeight)
	if err != nil {
		return nil, err
	}
	for _, txs := range res {
		for _, tx := range txs.Result.Txs {

			rawlog := strings.ReplaceAll(tx.TxResult.Log, `"{`, `{`)
			rawlog = strings.ReplaceAll(rawlog, `}"`, `}`)
			rawlog = strings.ReplaceAll(rawlog, `\n`, `,`)
			rawlog = strings.ReplaceAll(rawlog, `\`, ``)

			var logs []cosmostypes.RawLog
			err = json.Unmarshal([]byte(rawlog), &logs)
			if err!=nil{
				return nil,err
			}
			height, err := strconv.Atoi(tx.Height)
			if err != nil {
				return nil, err
			}

			balanceEntry, err := readlogs(logs, address, tx.Hash, height)
			if err != nil {
				return nil, err
			}

			balanceEntries = append(balanceEntries,
				balanceEntry...)
		}
	}
	accountBalanceSheet:=tabletypes.NewAccountBalanceSheet(address, balanceEntries)
	return	&accountBalanceSheet,nil
}

func readlogs(logs []cosmostypes.RawLog, address, hash string, height int) ([]tabletypes.BalanceEntry, error) {
	log.Trace().Str("module", "cosmos").Msg("reading logs")

	var balanceEntries []tabletypes.BalanceEntry

	for _, log := range logs {
		// There will be one transaction
		////fmt.Println(fmt.Sprintf("MsgIndex:%d", log.MsgIndex))
		in := "0"
		out := "0"
		msgType := ""

		// Read event for that log
		for _, event := range log.Events {
			//Catagorise each event and put it in a table
			attribute := ConvertAttributeToMap(event.Attributes)
			////fmt.Println(fmt.Sprintf("type:%s",event.Type))
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
			}
			if event.Type == "message" {
				bzaction, err := attribute["action"].MarshalJSON()
				if err != nil {
					return nil, err
				}
				msgType = strings.ReplaceAll(string(bzaction), "\"", "")
				////fmt.Println(fmt.Sprintf("action:%s", msgType))
			}
		}
		balanceEntries = append(balanceEntries,
			tabletypes.NewBalanceEntry(height, hash, in, out, msgType))
	}

	return balanceEntries, nil
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
	pageCount := math.MaxInt
	for page := 1; (lastHeight > targetHeight) && (pageCount >= page); page++ {
		query := fmt.Sprintf(`%s/tx_search?query="message.sender='%s'"&prove=true&page=%d&per_page=%d&order_by="desc"`,
			api, address, page, limit)

		log.Trace().Str("module", "cosmos").Str("Query tx:",query).Msg("Query from readTxs")

		resp, err := http.Get(query)
		if err != nil {
			return nil, fmt.Errorf("Fail to get tx from rpc:%s", err)
		}
		if resp.StatusCode != 200 {
			return nil, fmt.Errorf("Fail to get tx from rpc:Status :%s", resp.Status)
		}

		defer resp.Body.Close()

		bz, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		var txSearchRes cosmostypes.TxSearchRespond
		err = json.Unmarshal(bz, &txSearchRes)
		if err != nil {
			return nil, fmt.Errorf("Fail to marshal:%s", err)
		}
		if len(txSearchRes.Result.Txs) == 0 {
			return nil, nil
		}
		lastHeight, err = strconv.Atoi(
			txSearchRes.Result.Txs[len(txSearchRes.Result.Txs)-1].Height,
		)
		if err != nil {
			return nil, err
		}
		if pageCount == math.MaxInt {
			totalCount, err := strconv.Atoi(txSearchRes.Result.TotalCount)
			if err != nil {
				return nil, err
			}
			pageCount = totalCount/limit + 1
		}
		res = append(res, &txSearchRes)
	}

	return res, nil
}
