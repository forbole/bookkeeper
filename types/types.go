package types

import (
	"fmt"
	"time"
)

type Coin struct{
	CoinType string
	Quantity float64
	StartDate time.Time
}

func NewCoin(cointype string, quantity float64, startDate time.Time)Coin{
	return Coin{
		CoinType: cointype,
		Quantity: quantity,
		StartDate: startDate,
	}
}

type Balance struct{
	Coin *Coin
	VSCurrency string
	Price float32 // The coin price on that time
	Date time.Time // The time that the balance due
	Balance float32 // The balance on that account
}

func NewBalance(coin *Coin, price float32, date time.Time,balance float32,vSCurrency string)Balance{
	return Balance{
		Coin:coin,
		Price: price,
		Date:date,
		Balance:balance,
		VSCurrency: vSCurrency,
	}
}

type Balances []Balance

func (v Balances) GetCSV() string{
	outputcsv:="date,coin,VSCurrency coin price, account balance\n"
	for _,b:=range v{
		outputcsv+=fmt.Sprintf("%s,%s,%s,%f,%f\n",
			b.Date.String(),b.Coin.CoinType,b.VSCurrency,b.Price,b.Balance)
	}
	return outputcsv
}