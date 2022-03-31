package staking

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/forbole/bookkeeper/module/cosmos/types"
)

//https://api.cosmos.network/txs?message.module=distribution&transfer.sender=cosmos1m73mgwn3cm2e8x9a9axa0kw8nqz8a4927ywyqq

func BankSend(api string, address string, limit int, page int) error {
	resp, err := http.Get(fmt.Sprintf("%s/txs?message.module=distribution&transfer.sender=%s&limit=%d&page=%d",
		api, address, limit, page))
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	bz, err := ioutil.ReadAll(resp.Body)
	bz

	var bankreceipt types.BankRecipientRPC
	err = json.Unmarshal(bz, &bankreceipt)
	if err != nil {
		return err
	}

	return nil
}
