package types

import (
	"time"
)

type DateQuantity struct {
	Quantity float64
	Date     time.Time
}

func NewDateQuantity(quantity float64, date time.Time) DateQuantity {
	return DateQuantity{
		Quantity: quantity,
		Date:     date,
	}
}

// An coin history of how the balance in a account flow
type Coin struct {
	CoinType       string
	DateQuantities []DateQuantity
}

func NewCoin(cointype string, dateQuantities []DateQuantity) Coin {
	return Coin{
		CoinType:       cointype,
		DateQuantities: dateQuantities,
	}
}
