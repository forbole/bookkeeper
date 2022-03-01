package types

import "time"

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
	Coin Coin
	VSCurrency string
	Price uint64 // The price on that time
	Date time.Time // The time that the balance due
	Balance uint64 // The balance on that account
}

func NewBalance(coin Coin, price uint64, date time.Time,balance uint64,vSCurrency string)Balance{
	return Balance{
		Coin:coin,
		Price: price,
		Date:date,
		Balance:balance,
		VSCurrency: vSCurrency,
	}
}