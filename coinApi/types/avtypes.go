package avtypes

type ExchangeRate struct {
	RealtimeCurrencyExchangeRate struct {
		OneFromCurrencyCode string `json:"1. From_Currency Code"`
		TwoFromCurrencyName string `json:"2. From_Currency Name"`
		ThreeToCurrencyCode string `json:"3. To_Currency Code"`
		FourToCurrencyName  string `json:"4. To_Currency Name"`
		FiveExchangeRate    string `json:"5. Exchange Rate"`
		SixLastRefreshed    string `json:"6. Last Refreshed"`
		SevenTimeZone       string `json:"7. Time Zone"`
		EightBidPrice       string `json:"8. Bid Price"`
		NineAskPrice        string `json:"9. Ask Price"`
	} `json:"Realtime Currency Exchange Rate"`
}
