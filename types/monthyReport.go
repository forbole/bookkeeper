package types

import (
	"fmt"
	"math"
	"time"
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
	exponent := math.Pow(10, -6)

	for _, b := range v {
		outputcsv += fmt.Sprintf("%s,%s,%f,%f\n",
			b.FromDate,b.ToDate, float64(b.Commission)*exponent, float64(b.Reward)*exponent)
			commissionSum+=b.Commission
			rewardSum+=b.Reward
	}

	outputcsv+=fmt.Sprintf("\n Sum, ,%f,%f\n",float64(commissionSum)*exponent,float64(rewardSum)*exponent)
	
	return outputcsv
}