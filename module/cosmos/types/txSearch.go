package types

import "encoding/json"

type TxSearchRespond struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  struct {
		Txs []struct {
			Hash     string `json:"hash"`
			Height   string `json:"height"`
			Index    int    `json:"index"`
			TxResult struct {
				Code      int    `json:"code"`
				Data      string `json:"data"`
				Log       string `json:"log"`
				Info      string `json:"info"`
				GasWanted string `json:"gas_wanted"`
				GasUsed   string `json:"gas_used"`
				Events    []struct {
					Type       string `json:"type"`
					Attributes []struct {
						Key   string `json:"key"`
						Value string `json:"value"`
						Index bool   `json:"index"`
					} `json:"attributes"`
				} `json:"events"`
				Codespace string `json:"codespace"`
			} `json:"tx_result"`
			Tx    string `json:"tx"`
			Proof struct {
				RootHash string `json:"root_hash"`
				Data     string `json:"data"`
				Proof    struct {
					Total    string   `json:"total"`
					Index    string   `json:"index"`
					LeafHash string   `json:"leaf_hash"`
					Aunts    []string `json:"aunts"`
				} `json:"proof"`
			} `json:"proof"`
		} `json:"txs"`
		TotalCount string `json:"total_count"`
	} `json:"result"`
}


type RawLog struct {
	Events []Event `json:"events"`
	MsgIndex int `json:"msg_index,omitempty"`
}

type Event struct {
	Type       string `json:"type"`
	Attributes []Attributes `json:"attributes"`
}

type Attributes struct{
	Key   string `json:"key"`
	Value json.RawMessage `json:"value"`
}