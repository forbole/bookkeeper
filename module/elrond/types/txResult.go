package types

import "encoding/json"

type TxResult struct {
	TxHash        string          `json:"txHash"`
	GasLimit      int             `json:"gasLimit"`
	GasPrice      int             `json:"gasPrice"`
	GasUsed       int             `json:"gasUsed"`
	MiniBlockHash string          `json:"miniBlockHash"`
	Nonce         int             `json:"nonce"`
	Receiver      string          `json:"receiver"`
	ReceiverShard int64           `json:"receiverShard"`
	Round         int             `json:"round"`
	Sender        string          `json:"sender"`
	SenderShard   int             `json:"senderShard"`
	Signature     string          `json:"signature"`
	Status        string          `json:"status"`
	Value         string          `json:"value"`
	Fee           string          `json:"fee"`
	Timestamp     int             `json:"timestamp"`
	Data          string          `json:"data"`
	Function      string          `json:"function"`
	Action        json.RawMessage `json:"action,omitempty"`
	Results       []struct {
		Hash           string `json:"hash"`
		Timestamp      int    `json:"timestamp"`
		Nonce          int    `json:"nonce"`
		GasLimit       int    `json:"gasLimit"`
		GasPrice       int    `json:"gasPrice"`
		Value          string `json:"value"`
		Sender         string `json:"sender"`
		Receiver       string `json:"receiver"`
		PrevTxHash     string `json:"prevTxHash"`
		OriginalTxHash string `json:"originalTxHash"`
		CallType       string `json:"callType"`
		MiniBlockHash  string `json:"miniBlockHash"`
		Data           string `json:"data,omitempty"`
		Logs           struct {
			Address string `json:"address"`
			Events  []struct {
				Address    string   `json:"address"`
				Identifier string   `json:"identifier"`
				Topics     []string `json:"topics"`
				Order      int      `json:"order"`
			} `json:"events"`
		} `json:"logs,omitempty"`
	} `json:"results"`
	Price float64 `json:"price"`
	Logs  struct {
		Address string          `json:"address"`
		Events  json.RawMessage `json:"events,omitempty"`
	} `json:"logs,omitempty"`
	Operations []struct {
		ID       string `json:"id"`
		Action   string `json:"action"`
		Type     string `json:"type"`
		Sender   string `json:"sender"`
		Receiver string `json:"receiver"`
		Value    string `json:"value"`
	} `json:"operations"`
}
