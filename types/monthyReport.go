package types

import (
	"fmt"
	"time"
)

// MonthyReportRow represent a row of monthy report
type MonthyReportRow struct{
	Date time.Time
	Commission int
	Reward int
}

func NewMonthyReportRow(date time.Time,commission int,reward int)MonthyReportRow{
	return MonthyReportRow{
		Date: date,
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

func (v MonthyReportRows) GetCSV()string{
	outputcsv := "height,txHash,Commission,Delegator_Reward\n"
	commissionSum:=0
	rewardSum:=0
	for _, b := range v {
		outputcsv += fmt.Sprintf("%s,%d,%d\n",
			b.Date, b.Reward, b.Commission)
			commissionSum+=b.Commission
			rewardSum+=b.Reward
	}
	outputcsv+=fmt.Sprintf("\n Sum, %d,%d\n",commissionSum,rewardSum)
	
	return outputcsv
}