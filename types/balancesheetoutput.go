package types

import (
	"fmt"
	"time"
)

// This represent a row in .csv file
type Balance struct {
	Coin       string
	VSCurrency string
	Quantity float64
	Price      float32   // The coin price on that time
	Date       time.Time // The time that the balance due
	Balance    float32   // The balance on that account
}

func NewBalance(coin string, price float32, date time.Time, balance float32, vSCurrency string,quantity float64) Balance {
	return Balance{
		Coin:       coin,
		Price:      price,
		Date:       date,
		Balance:    balance,
		VSCurrency: vSCurrency,
		Quantity: quantity,
	}
}

type Balances []Balance

func (v Balances) GetCSV() string {
	outputcsv := "date,coin,VSCurrency,quantity, coin price, account balance\n"
	for _, b := range v {
		outputcsv += fmt.Sprintf("%s,%s,%s,%f,%f,%f\n",
			b.Date.String(), b.Coin, b.VSCurrency, b.Quantity, b.Price, b.Balance)
	}
	return outputcsv
}

// Coin price and quantity in one entry
type CoinDetail struct{
	Coin string
	Price float32
	Quantity float64
	Balance float32
}

// It represent a row in Total Balance balance sheet for multiple type of coins
type TotalBalance struct{
	CoinDetails []CoinDetail
	VSCurrency string
	Date       time.Time // The time that the balance due
	TotalBalance    float32   // The balance on that account
}

type TotalBalances []TotalBalance

func (v TotalBalances) GetCSV() string{
	coinSchema:=""
	for i,_ :=range v[0].CoinDetails{
		coinSchema+=fmt.Sprintf("Coin %d, Quantity %d, Price %d, Balance%d,",i,i,i,i)
	}
	csv := fmt.Sprintf("date,%sVSCurrency,Total Balance\n",coinSchema)

	for _,totalBalance:=range v{
		coinDetailList:=""
		for _,coinDetail:=range totalBalance.CoinDetails{
			coinDetailList+=fmt.Sprintf("%s,%f,%f,%f,",coinDetail.Coin,coinDetail.Quantity,coinDetail.Price,coinDetail.Balance)
		}
		csv+=fmt.Sprintf("%s,%s%s,%f\n",totalBalance.Date,coinDetailList,totalBalance.VSCurrency,totalBalance.TotalBalance)
	}
	return csv
}

func BalanceToCoinDetail(balance Balance)CoinDetail{
	return CoinDetail{
		Coin: balance.Coin,
		Price: balance.Price,
		Quantity: balance.Quantity,
		Balance: balance.Balance,
	}
}

func NewTotalBalance(balance []Balance,totalBalance float32,date time.Time)TotalBalance{
	coinDetails := make([]CoinDetail,len(balance))
	for i,b:=range balance{
		coinDetails[i]=BalanceToCoinDetail(b)
	}
	return TotalBalance{
		CoinDetails: coinDetails,
		VSCurrency:balance[0].VSCurrency,
		Date:date,
		TotalBalance: totalBalance,
	}
}