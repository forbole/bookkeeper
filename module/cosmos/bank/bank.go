package bank

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/forbole/bookkeeper/module/cosmos/types"
)

//endpoint that can query all the txs
//https://rpc.cosmos.network/tx_search?query=%22message.sender=%27cosmos15mj8w79uf7gyxr7mnejz9k57ykcp4lc3mz3wly%27%22

func BankRecipient(api string, address string, limit int, page int) error {
	resp, err := http.Get(fmt.Sprintf("%s/txs?message.module=bank&transfer.recipient=%s&limit=%d&page=%d",
		api, address, limit, page))
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	bz, err := ioutil.ReadAll(resp.Body)

	var bankreceipt types.BankRecipientRPC
	err = json.Unmarshal(bz, &bankreceipt)
	if err != nil {
		return err
	}

	return nil
}

func BankSend(api string, address string, limit int, page int) error {
	resp, err := http.Get(fmt.Sprintf("%s/txs?message.module=bank&transfer.sender=%s&limit=%d&page=%d",
		api, address, limit, page))
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	bz, err := ioutil.ReadAll(resp.Body)

	var bankreceipt types.BankRecipientRPC
	err = json.Unmarshal(bz, &bankreceipt)
	if err != nil {
		return err
	}

	return nil
}
