package tabletypes

import (
	"fmt"
	"math/big"
	"time"
)

type DateValueRow struct{
	Date time.Time
	Delegation *big.Float
	Denom string
	DelegationValue *big.Float
	TxHash string
	Types string
}

func NewDateValueRow(date time.Time,delegation *big.Float,denom string,
	delegationValue *big.Float,txHash string, types string)DateValueRow{
		return DateValueRow{
			Date: date,
			Delegation: delegation,
			Denom: denom,
			DelegationValue: delegationValue,
			TxHash: txHash,
			Types: types,
		}
}

type DateValueRows []DateValueRow

func (v DateValueRows)GetCSV()string{
	csv:="Date,Delegation,Denom,DelegationValue,TxHash,Types\n"
	for _,row:=range v{
		csv+=fmt.Sprintf("%s,%f,%s,%f,%s,%s\n",row.Date.String(),row.Delegation,
		row.Denom,row.DelegationValue,row.TxHash,row.Types)
	}
	return csv
}