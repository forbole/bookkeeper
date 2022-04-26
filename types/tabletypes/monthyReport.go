package tabletypes

import (
	"fmt"
	"math"
	"math/big"

	//"net/http"
	"time"

	//coingecko "github.com/superoo7/go-gecko/v3"
	//"github.com/forbole/bookkeeper/utils"
	"github.com/forbole/bookkeeper/types"
	"github.com/forbole/bookkeeper/utils"
)

// MonthyReportRow represent a row of monthy report
type MonthyReportRow struct{
	FromDate time.Time
	ToDate time.Time

	Commission *big.Int
	Reward *big.Int
	Denom string
}

func NewMonthyReportRow(fromDate time.Time,toDate time.Time,commission *big.Int,reward *big.Int,denom string)MonthyReportRow{
	return MonthyReportRow{
		FromDate: fromDate,
		ToDate: toDate,

		Commission: commission,
		Reward: reward,
		Denom: denom,
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
	outputcsv := "From_date,to_date,Commission,Delegator_Reward,denom\n"
	rewardSum:=big.NewInt(0)
	commissionSum:=big.NewInt(0)

	exponent:=new(big.Float).SetFloat64((math.Pow(10,float64(-1 * exp))))
	//exponent := math.Pow(10, float64(-1 * exp))

	for _, b := range v {
		// Change to big float with exp
		c:=new(big.Float).SetInt(b.Commission)
		r:=new(big.Float).SetInt(b.Reward)

		outputcsv += fmt.Sprintf("%s,%s,%v,%v,%s\n",
			b.FromDate,b.ToDate, c.Mul(c,exponent),r.Mul(r,exponent),b.Denom)

			commissionSum.Add(commissionSum,b.Commission)
			rewardSum.Add(rewardSum,b.Reward)
	}

	cs:=new(big.Float).SetInt(commissionSum)
	rs:=new(big.Float).SetInt(rewardSum)

	outputcsv+=fmt.Sprintf("\n Sum, ,%v,%v\n",cs.Mul(cs,exponent),rs.Mul(rs,exponent))
	
	return outputcsv
}

// GetCSV generate the monthy report and turn the result into exponent form
func (v MonthyReportRows) GetCSVConvertedPrice(denom []types.Denom, vsCurrency string)(string,error){	
	
	rewardSum:=big.NewInt(0)
	commissionSum:=big.NewInt(0)
	denomMap,err:=utils.ConvertDenomToMap(denom)
	if err!=nil{
		return "",err
	}
	
	/*
	outputcsv := "From_date,to_date,Commission,Delegator_Reward, ,Commission_value ,Delegator_Reward_value\n"

	exponent:=new(big.Float).SetFloat64((math.Pow(10,float64(-1 * exp))))

	for _, b := range v {
		c:=new(big.Float).SetInt(b.Commission)
		r:=new(big.Float).SetInt(b.Reward)
		cexp:=new(big.Float).Mul(c,exponent)
		rexp:=new(big.Float).Mul(r,exponent)
		cexpCoinPrice:=new(big.Float).Mul(cexp,coinprice)
		rexpCoinPrice:=new(big.Float).Mul(rexp,coinprice)
		outputcsv += fmt.Sprintf("%s,%s,%f,%f, ,%f,%f\n",
			b.FromDate,b.ToDate, cexp,rexp,cexpCoinPrice,rexpCoinPrice)
			commissionSum.Add(commissionSum,b.Commission)
			rewardSum.Add(rewardSum,b.Reward)
	}

	cs:=new(big.Float).SetInt(commissionSum)
	rs:=new(big.Float).SetInt(rewardSum)

	realCommissionSum:=new(big.Float).Mul(cs,exponent)
	realRewardSum:=new(big.Float).Mul(rs,exponent)
	
	outputcsv+=fmt.Sprintf("\n Sum, ,%f,%f, ,%f,%f\n",
	realCommissionSum,realRewardSum,
	new(big.Float).Mul(realCommissionSum,coinprice) ,new(big.Float).Mul(realRewardSum,coinprice))
	 */
	outputcsv:="Not implemented"
	return outputcsv,nil
}
