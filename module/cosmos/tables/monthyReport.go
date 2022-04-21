package cosmos

import (
	"fmt"
	"time"

	"github.com/forbole/bookkeeper/module/cosmos/utils"
	"github.com/forbole/bookkeeper/types"
)

// GetMonthyReport get monthy report from now to desired date(in the past)
func GetMonthyReport(details types.IndividualChain,from time.Time)([]types.AddressMonthyReport,error){
	var monthyReports []types.AddressMonthyReport

	balanceEntries,err:=GetTxs(details)
	if err!=nil{
		return nil,err
	}
	t:=*(lastMonth(time.Now()))
	for _,b:=range balanceEntries{
		var monthyReportRows types.MonthyReportRows
		rewardCommission,err:=GetRewardCommission(b,details.Denom)
		if err!=nil{
			return nil,err
		}
		fmt.Println(rewardCommission.Rows.GetCSV())
		i:=0
		for t.After(from){
			targetHeight,err:=utils.GetHeightByDate(t,details.LcdEndpoint)
			if err!=nil{
				return nil,err
			}
			commission:=0
			reward:=0
			rows:=rewardCommission.Rows
			for rows[i].Height>targetHeight && i<len(rows)-1{
					commission+=rows[i].Commission
					reward+=rows[i].Reward
					i++
			}
			fmt.Println(reward)
			fmt.Println(commission)
			monthyReportRows=append(monthyReportRows,types.NewMonthyReportRow(t,*(nextMonth(t)),commission,reward))
			t=*(lastMonth(t))
	
		}
		monthyReports=append(monthyReports,types.NewAddressMonthyReport(b.Address,monthyReportRows))
	}
	return monthyReports,nil
}

func nextMonth(t time.Time)*time.Time{
	date:=t

	day:=t.Day()
	if t.Day()>28{
		day=28
	}

	switch t.Month(){
	case time.January:
		date=time.Date(t.Year(),time.February,day,t.Hour(),t.Minute(),t.Second(),t.Nanosecond(),t.Location())
		break
	case time.February:
		date=time.Date(t.Year(),time.March,day,t.Hour(),t.Minute(),t.Second(),t.Nanosecond(),t.Location())
		break
	case time.March:
		date=time.Date(t.Year(),time.April,day,t.Hour(),t.Minute(),t.Second(),t.Nanosecond(),t.Location())
		break
	case time.April:
		date=time.Date(t.Year(),time.May,day,t.Hour(),t.Minute(),t.Second(),t.Nanosecond(),t.Location())
		break
	case time.May:
		date=time.Date(t.Year(),time.June,day,t.Hour(),t.Minute(),t.Second(),t.Nanosecond(),t.Location())
		break
	case time.June:
		date=time.Date(t.Year(),time.July,day,t.Hour(),t.Minute(),t.Second(),t.Nanosecond(),t.Location())
		break
	case time.July:
		date=time.Date(t.Year(),time.August,day,t.Hour(),t.Minute(),t.Second(),t.Nanosecond(),t.Location())
		break
	case time.August:
		date=time.Date(t.Year(),time.September,day,t.Hour(),t.Minute(),t.Second(),t.Nanosecond(),t.Location())
		break
	case time.September:
		date=time.Date(t.Year(),time.October,day,t.Hour(),t.Minute(),t.Second(),t.Nanosecond(),t.Location())
		break
	case time.October:
		date=time.Date(t.Year(),time.November,day,t.Hour(),t.Minute(),t.Second(),t.Nanosecond(),t.Location())
		break
	case time.November:
		date=time.Date(t.Year(),time.December,day,t.Hour(),t.Minute(),t.Second(),t.Nanosecond(),t.Location())
		break
	case time.December:
		date=time.Date(t.Year()+1,time.January,day,t.Hour(),t.Minute(),t.Second(),t.Nanosecond(),t.Location())
		break
	} 

	return &date
}

func lastMonth(t time.Time)*time.Time{
	date:=t

	day:=t.Day()
	if t.Day()>28{
		day=28
	}

	switch t.Month(){
	case time.January:
		date=time.Date(t.Year()-1,time.December,day,t.Hour(),t.Minute(),t.Second(),t.Nanosecond(),t.Location())
		break
	case time.February:
		date=time.Date(t.Year(),time.January,day,t.Hour(),t.Minute(),t.Second(),t.Nanosecond(),t.Location())
		break
	case time.March:
		date=time.Date(t.Year(),time.February,day,t.Hour(),t.Minute(),t.Second(),t.Nanosecond(),t.Location())
		break
	case time.April:
		date=time.Date(t.Year(),time.March,day,t.Hour(),t.Minute(),t.Second(),t.Nanosecond(),t.Location())
		break
	case time.May:
		date=time.Date(t.Year(),time.April,day,t.Hour(),t.Minute(),t.Second(),t.Nanosecond(),t.Location())
		break
	case time.June:
		date=time.Date(t.Year(),time.May,day,t.Hour(),t.Minute(),t.Second(),t.Nanosecond(),t.Location())
		break
	case time.July:
		date=time.Date(t.Year(),time.June,day,t.Hour(),t.Minute(),t.Second(),t.Nanosecond(),t.Location())
		break
	case time.August:
		date=time.Date(t.Year(),time.July,day,t.Hour(),t.Minute(),t.Second(),t.Nanosecond(),t.Location())
		break
	case time.September:
		date=time.Date(t.Year(),time.August,day,t.Hour(),t.Minute(),t.Second(),t.Nanosecond(),t.Location())
		break
	case time.October:
		date=time.Date(t.Year(),time.September,day,t.Hour(),t.Minute(),t.Second(),t.Nanosecond(),t.Location())
		break
	case time.November:
		date=time.Date(t.Year(),time.October,day,t.Hour(),t.Minute(),t.Second(),t.Nanosecond(),t.Location())
		break
	case time.December:
		date=time.Date(t.Year(),time.November,day,t.Hour(),t.Minute(),t.Second(),t.Nanosecond(),t.Location())
		break
	} 

	return &date
}

