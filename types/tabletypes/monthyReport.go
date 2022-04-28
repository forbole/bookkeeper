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
	
	rewardSum:=big.NewFloat(0)
	commissionSum:=big.NewFloat(0)
	denomMap,err:=utils.ConvertDenomToMap(denom,vsCurrency)
	if err!=nil{
		return "",err
	}

	outputcsv := "From_date,to_date,Commission_value ,Delegator_Reward_value\n"

	currentFromDate:=v[0].FromDate
	rewardInMonth:=new(big.Float).SetFloat64(0)
	commissionInMonth:=new(big.Float).SetFloat64(0)
	for i,row:=range v{
			// since they are same day, add it together
			if _, ok := denomMap[row.Denom]; !ok {
				// skip if that is not exist
				fmt.Println(fmt.Sprintf("Coin is not supported:%s",row.Denom))
				continue
			}
			c:=new(big.Float).SetInt(row.Commission)
			commission:=new(big.Float).Mul(c,denomMap[row.Denom].Exponent)
			commissionConverted:=new(big.Float).Mul(commission,denomMap[row.Denom].Price)
			newCommission:=new(big.Float).Add(commissionConverted,commissionInMonth)
			commissionInMonth=newCommission

			r:=new(big.Float).SetInt(row.Reward)
			reward:=new(big.Float).Mul(r,denomMap[row.Denom].Exponent)
			rewardConverted:=new(big.Float).Mul(reward,denomMap[row.Denom].Price)
			newRewardInMonth:=new(big.Float).Add(rewardConverted,rewardInMonth)
			rewardInMonth=newRewardInMonth
			fmt.Println(rewardInMonth)
			fmt.Println(commissionInMonth)

		
		if i+1==len(v) || v[i+1].FromDate!=currentFromDate{
			// If next entry changed date, write
			outputcsv += fmt.Sprintf("%s,%s,%f,%f,\n",
			row.FromDate,row.ToDate,rewardInMonth,commissionInMonth)
			newRewardSum:=new(big.Float).Add(rewardSum,rewardInMonth)
			newCommissionSum:=new(big.Float).Add(commissionInMonth,commissionSum)

			rewardSum=newRewardSum
			commissionSum=newCommissionSum
			if i+1==len(v){
				break
			}


			// Date changed, reset
			rewardInMonth=new(big.Float).SetFloat64(0)
			commissionInMonth=new(big.Float).SetFloat64(0)
			currentFromDate=v[i+1].FromDate
		}
	}

	outputcsv+=fmt.Sprintf("sum,,%f,%f",rewardSum,commissionSum)
	
	return outputcsv,nil
}