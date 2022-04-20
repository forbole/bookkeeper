package types

import (
	"fmt"
	"math/big"
	"math"
	"net/http"
	"time"

	coingecko "github.com/superoo7/go-gecko/v3"
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

// GetCSV generate the monthy report and turn the result into exponent form
func (v MonthyReportRows) GetCSVConvertedPrice(exp int, coinId string, vsCurrency string)(string,error){
	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}

	rewardSum:=big.NewInt(0)
	commissionSum:=big.NewInt(0)
	
	CG := coingecko.NewClient(httpClient)

	singlePrice,err:=CG.CoinsMarket(vsCurrency, []string{coinId}, "", 0, 0, false, nil)
	if err!=nil{
		return "",err
	}
	fmt.Println((*singlePrice)[0].CurrentPrice)
	c:=float64((*singlePrice)[0].CurrentPrice)
	fmt.Println(c)
	coinprice:=new(big.Float).SetFloat64(c)

	outputcsv := "From_date,to_date,Commission,Delegator_Reward, ,Commission_Converted ,Delegator_Reward_Converted\n"

	exponent:=new(big.Float).SetFloat64((math.Pow(10,float64(-1 * exp))))

	for _, b := range v {
		c:=new(big.Float).SetInt(b.Commission)
		r:=new(big.Float).SetInt(b.Reward)
		cexp:=c.Mul(c,exponent)
		rexp:=r.Mul(r,exponent)
		outputcsv += fmt.Sprintf("%s,%s,%f,%f, ,%f,%f\n",
			b.FromDate,b.ToDate, cexp,rexp,cexp.Mul(cexp,coinprice),rexp.Mul(rexp,coinprice))
			commissionSum.Add(commissionSum,b.Commission)
			rewardSum.Add(rewardSum,b.Reward)
	}

	cs:=new(big.Float).SetInt(commissionSum)
	rs:=new(big.Float).SetInt(rewardSum)

	realCommissionSum:=cs.Mul(cs,exponent)
	realRewardSum:=rs.Mul(rs,exponent)
	
	outputcsv+=fmt.Sprintf("\n Sum, ,%f,%f, ,%f,%f\n",
	realCommissionSum,realRewardSum,
	realCommissionSum.Mul(realCommissionSum,coinprice) ,realRewardSum.Mul(realRewardSum,coinprice))
	
	return outputcsv,nil
}