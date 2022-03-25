package types

import (
	"time"
)

type BankRecipientRPC struct {
	TotalCount string `json:"total_count"`
	Count      string `json:"count"`
	PageNumber string `json:"page_number"`
	PageTotal  string `json:"page_total"`
	Limit      string `json:"limit"`
	Txs        []struct {
		Height string `json:"height"`
		Txhash string `json:"txhash"`
		Data   string `json:"data"`
		RawLog string `json:"raw_log"`
		Logs   []struct {
			Events []struct {
				Type       string `json:"type"`
				Attributes []struct {
					Key   string `json:"key"`
					Value string `json:"value"`
				} `json:"attributes"`
			} `json:"events"`
		} `json:"logs"`
		GasWanted string `json:"gas_wanted"`
		GasUsed   string `json:"gas_used"`
		Tx        struct {
			Type  string `json:"type"`
			Value struct {
				Msg []struct {
					Type  string `json:"type"`
					Value struct {
						FromAddress string `json:"from_address"`
						ToAddress   string `json:"to_address"`
						Amount      []struct {
							Denom  string `json:"denom"`
							Amount string `json:"amount"`
						} `json:"amount"`
					} `json:"value"`
				} `json:"msg"`
				Fee struct {
					Amount []struct {
						Denom  string `json:"denom"`
						Amount string `json:"amount"`
					} `json:"amount"`
					Gas string `json:"gas"`
				} `json:"fee"`
				Signatures    []interface{} `json:"signatures"`
				Memo          string        `json:"memo"`
				TimeoutHeight string        `json:"timeout_height"`
			} `json:"value"`
		} `json:"tx"`
		Timestamp time.Time `json:"timestamp"`
	} `json:"txs"`
}