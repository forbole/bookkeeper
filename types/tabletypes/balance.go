package tabletypes

import (
	"fmt"
)

// BalanceEntry represent a row of csv
// This is raw tx data for a address
type BalanceEntry struct {
	Height  int
	TxHash  string
	In      string
	Out     string
	MsgType string
}

func NewBalanceEntry(height int, txHash string, in string, out string, msgType string) BalanceEntry {
	return BalanceEntry{
		Height:  height,
		TxHash:  txHash,
		In:      in,
		Out:     out,
		MsgType: msgType,
	}
}

type BalanceEntries []BalanceEntry

func (v BalanceEntries) GetCSV() string {
	outputcsv := "height,txHash,receive_amount,sent_amount, msgType\n"
	for _, b := range v {
		outputcsv += fmt.Sprintf("%d,%s,%s,%s,%s\n",
			b.Height, b.TxHash, b.In, b.Out, b.MsgType)
	}
	return outputcsv
}

type AddressBalanceEntry struct {
	Address string
	Rows    BalanceEntries
}

func NewAccountBalanceSheet(address string, balanceEntry []BalanceEntry) AddressBalanceEntry {
	return AddressBalanceEntry{
		Address: address,
		Rows:    balanceEntry,
	}
}
