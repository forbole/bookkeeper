package types

import "encoding/json"

type Tx struct {
	TxHash        string `json:"txHash"`
	GasLimit      int    `json:"gasLimit"`
	GasPrice      int    `json:"gasPrice"`
	GasUsed       int    `json:"gasUsed"`
	MiniBlockHash string `json:"miniBlockHash"`
	Nonce         int    `json:"nonce"`
	Receiver      string `json:"receiver"`
	ReceiverShard int64  `json:"receiverShard"`
	Round         int    `json:"round"`
	Sender        string `json:"sender"`
	SenderShard   int    `json:"senderShard"`
	Signature     string `json:"signature"`
	Status        string `json:"status"`
	Value         string `json:"value"`
	Fee           string `json:"fee"`
	Timestamp     int    `json:"timestamp"`
	Data          string `json:"data"`
	Function      string `json:"function"`
	Action        json.RawMessage `json:"action,omitmpty"`
}