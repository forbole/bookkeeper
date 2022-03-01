package types

import "time"

type Coin struct{
	CoinType string
	Address string
	StartDate time.Time
}

func NewCoin(cointype string, address string, startDate time.Time)Coin{
	return Coin{
		CoinType: cointype,
		Address: address,
		StartDate: startDate,
	}
}

type Balance struct{
	Coin Coin
	Price uint64 // The price on that time
	Date time.Time // The time that the balance due
	Balance uint64 // The balance on that account
}

func NewBalance(coin Coin, price uint64, date time.Time,balance uint64)Balance{
	return Balance{
		Coin:coin,
		Price: price,
		Date:date,
		Balance:balance,
	}
}