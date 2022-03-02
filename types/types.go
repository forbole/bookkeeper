package types

import (
	"fmt"
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

// This represent a row in .csv file
type Balance struct {
	Coin       string
	VSCurrency string
	Price      float32   // The coin price on that time
	Date       time.Time // The time that the balance due
	Balance    float32   // The balance on that account
}

func NewBalance(coin string, price float32, date time.Time, balance float32, vSCurrency string) Balance {
	return Balance{
		Coin:       coin,
		Price:      price,
		Date:       date,
		Balance:    balance,
		VSCurrency: vSCurrency,
	}
}

type Balances []Balance

func (v Balances) GetCSV() string {
	outputcsv := "date,coin,VSCurrency, coin price, account balance\n"
	for _, b := range v {
		outputcsv += fmt.Sprintf("%s,%s,%s,%f,%f\n",
			b.Date.String(), b.Coin, b.VSCurrency, b.Price, b.Balance)
	}
	return outputcsv
}
