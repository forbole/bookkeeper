package types

import (
	"fmt"
	"math/big"
	"math"
	"time"
)

// MonthyReportRow represent a row of monthy report
type MonthyReportRow struct{
	FromDate time.Time
	ToDate time.Time

	Commission *big.Int
	Reward *big.Int
}

func NewMonthyReportRow(fromDate time.Time,toDate time.Time,commission *big.Int,reward *big.Int)MonthyReportRow{
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
	rewardSum:=big.NewInt(0)
	commissionSum:=big.NewInt(0)

	exponent:=new(big.Float).SetFloat64((math.Pow(10,float64(-1 * exp))))
	//exponent := math.Pow(10, float64(-1 * exp))

	for _, b := range v {
		// Change to big float with exp
		c:=new(big.Float).SetInt(b.Commission)
		r:=new(big.Float).SetInt(b.Reward)

		outputcsv += fmt.Sprintf("%s,%s,%v,%v\n",
			b.FromDate,b.ToDate, c.Mul(c,exponent),r.Mul(r,exponent))

			commissionSum.Add(commissionSum,b.Commission)
			rewardSum.Add(rewardSum,b.Reward)
	}

	cs:=new(big.Float).SetInt(commissionSum)
	rs:=new(big.Float).SetInt(rewardSum)

	outputcsv+=fmt.Sprintf("\n Sum, ,%v,%v\n",cs.Mul(cs,exponent),rs.Mul(rs,exponent))
	
	return outputcsv
}