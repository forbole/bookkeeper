package types

import (
	"fmt"
	"math"
	"net/http"
	"time"

	coingecko "github.com/superoo7/go-gecko/v3"
)

// MonthyReportRow represent a row of monthy report
type MonthyReportRow struct{
	FromDate time.Time
	ToDate time.Time

	Commission int
	Reward int
}

func NewMonthyReportRow(fromDate time.Time,toDate time.Time,commission int,reward int)MonthyReportRow{
	return MonthyReportRow{
		FromDate: fromDate,
		ToDate: toDate,

		Commission: commission,
		Reward: reward,
	}
}

type MonthyReportRows []MonthyReportRow


type AddressMonthyReport struct{
	Address string
	Rows MonthyReportRows
}

func NewAddressMonthyReport(address string,rewardCommissions MonthyReportRows)AddressMonthyReport{
	return AddressMonthyReport{
		Address: address,
		Rows: rewardCommissions,
	}
}

// GetCSV generate the monthy report and turn the result into exponent form
func (v MonthyReportRows) GetCSV(exp int)string{
	outputcsv := "From_date,to_date,Commission,Delegator_Reward\n"
	commissionSum:=0
	rewardSum:=0
	exponent := math.Pow(10, float64(-1*exp))

	for _, b := range v {
		outputcsv += fmt.Sprintf("%s,%s,%f,%f\n",
			b.FromDate,b.ToDate, float64(b.Commission)*exponent, float64(b.Reward)*exponent)
			commissionSum+=b.Commission
			rewardSum+=b.Reward
	}

	outputcsv+=fmt.Sprintf("\n Sum, ,%f,%f\n",float64(commissionSum)*exponent,float64(rewardSum)*exponent)
	
	return outputcsv
}

// GetCSV generate the monthy report and turn the result into exponent form
func (v MonthyReportRows) GetCSVConvertedPrice(exp int, coinId string, vsCurrency string)(string,error){
	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}
	
	CG := coingecko.NewClient(httpClient)

	singlePrice,err:=CG.CoinsMarket(vsCurrency, []string{coinId}, "", 0, 0, false, nil)
	if err!=nil{
		return "",err
	}
	if len(*singlePrice)==0{
		return "",fmt.Errorf("Error getting coinsmarket")
	}
	fmt.Println((*singlePrice)[0].CurrentPrice)
	coinprice:=float64((*singlePrice)[0].CurrentPrice)
	fmt.Println(coinprice)

	outputcsv := "From_date,to_date,Commission,Delegator_Reward, ,Commission_Converted ,Delegator_Reward_Converted\n"
	commissionSum:=0
	rewardSum:=0
	exponent := math.Pow(10, float64(-1*exp))

	for _, b := range v {
		c:=float64(b.Commission)*exponent
		r:=float64(b.Reward)*exponent
		outputcsv += fmt.Sprintf("%s,%s,%f,%f, ,%f,%f\n",
			b.FromDate,b.ToDate, c, r,c*coinprice,r*coinprice)
			commissionSum+=b.Commission
			rewardSum+=b.Reward
	}

	realCommissionSum:=float64(commissionSum)*exponent
	realRewardSum:=float64(rewardSum)*exponent
	
	outputcsv+=fmt.Sprintf("\n Sum, ,%f,%f, ,%f,%f\n",
	realCommissionSum,realRewardSum,realCommissionSum*coinprice ,realRewardSum*coinprice)
	
	return outputcsv,nil
}