package cosmos

import (
	"time"

	"github.com/forbole/bookkeeper/types"
	"github.com/forbole/bookkeeper/module/cosmos/utils"

)

// GetMonthyReport get monthy report from now to desired date(in the past)
func GetMonthyReport(details types.IndividualChain,from time.Time)(types.MonthyReportRows,error){
	var monthyReport types.MonthyReportRows

	balanceEntries,err:=GetTxs(details)
	if err!=nil{
		return nil,err
	}
	t:=from
	for _,b:=range balanceEntries{
		rewardCommission,err:=GetRewardCommission(b,"uatom")
		if err!=nil{
			return nil,err
		}
		i:=0
		for time.Now().After(t){
			targetHeight,err:=utils.GetHeightFormDate(t,details.LcdEndpoint)
			if err!=nil{
				return nil,err
			}
			commission:=0
			reward:=0
			rows:=rewardCommission.Rows
			for rows[i].Height<targetHeight{
					commission+=rows[i].Commission
					reward+=rows[i].Commission
					i++
			}
			monthyReport=append(monthyReport,types.NewMonthyReportRow(t,commission,reward))
			t=*(lastMonth(t))
		}

	}
}

func nextMonth(t time.Time)*time.Time{
	date:=t
	switch t.Month(){
	case time.January:
		day:=t.Day()
		if t.Day()>28{
			day=28
		}
		date=time.Date(t.Year(),time.February,day,t.Hour(),t.Minute(),t.Second(),t.Nanosecond(),t.Location())
		break
	case time.February:
		date=time.Date(t.Year(),time.March,bigMonthConvert(t.Day()),t.Hour(),t.Minute(),t.Second(),t.Nanosecond(),t.Location())
		break
	case time.March:
		date=time.Date(t.Year(),time.April,bigMonthConvert(t.Day()),t.Hour(),t.Minute(),t.Second(),t.Nanosecond(),t.Location())
		break
	case time.April:
		date=time.Date(t.Year(),time.May,bigMonthConvert(t.Day()),t.Hour(),t.Minute(),t.Second(),t.Nanosecond(),t.Location())
		break
	case time.May:
		date=time.Date(t.Year(),time.June,bigMonthConvert(t.Day()),t.Hour(),t.Minute(),t.Second(),t.Nanosecond(),t.Location())
		break
	case time.June:
		date=time.Date(t.Year(),time.July,bigMonthConvert(t.Day()),t.Hour(),t.Minute(),t.Second(),t.Nanosecond(),t.Location())
		break
	case time.July:
		date=time.Date(t.Year(),time.August,bigMonthConvert(t.Day()),t.Hour(),t.Minute(),t.Second(),t.Nanosecond(),t.Location())
		break
	case time.August:
		date=time.Date(t.Year(),time.September,bigMonthConvert(t.Day()),t.Hour(),t.Minute(),t.Second(),t.Nanosecond(),t.Location())
		break
	case time.September:
		date=time.Date(t.Year(),time.October,bigMonthConvert(t.Day()),t.Hour(),t.Minute(),t.Second(),t.Nanosecond(),t.Location())
		break
	case time.October:
		date=time.Date(t.Year(),time.November,bigMonthConvert(t.Day()),t.Hour(),t.Minute(),t.Second(),t.Nanosecond(),t.Location())
		break
	case time.November:
		date=time.Date(t.Year(),time.December,bigMonthConvert(t.Day()),t.Hour(),t.Minute(),t.Second(),t.Nanosecond(),t.Location())
		break
	case time.December:
		date=time.Date(t.Year()+1,time.January,bigMonthConvert(t.Day()),t.Hour(),t.Minute(),t.Second(),t.Nanosecond(),t.Location())
		break
	} 

	return &date
}

func lastMonth(t time.Time)*time.Time{
	date:=t
	switch t.Month(){
	case time.January:
		day:=t.Day()
		if t.Day()>28{
			day=28
		}
		date=time.Date(t.Year()-1,time.December,day,t.Hour(),t.Minute(),t.Second(),t.Nanosecond(),t.Location())
		break
	case time.February:
		date=time.Date(t.Year(),time.January,bigMonthConvert(t.Day()),t.Hour(),t.Minute(),t.Second(),t.Nanosecond(),t.Location())
		break
	case time.March:
		date=time.Date(t.Year(),time.February,bigMonthConvert(t.Day()),t.Hour(),t.Minute(),t.Second(),t.Nanosecond(),t.Location())
		break
	case time.April:
		date=time.Date(t.Year(),time.March,bigMonthConvert(t.Day()),t.Hour(),t.Minute(),t.Second(),t.Nanosecond(),t.Location())

		break
	case time.May:
		date=time.Date(t.Year(),time.April,bigMonthConvert(t.Day()),t.Hour(),t.Minute(),t.Second(),t.Nanosecond(),t.Location())

		break
	case time.June:
		date=time.Date(t.Year(),time.May,bigMonthConvert(t.Day()),t.Hour(),t.Minute(),t.Second(),t.Nanosecond(),t.Location())
		break
	case time.July:
		date=time.Date(t.Year(),time.June,bigMonthConvert(t.Day()),t.Hour(),t.Minute(),t.Second(),t.Nanosecond(),t.Location())

		break
	case time.August:
		date=time.Date(t.Year(),time.July,bigMonthConvert(t.Day()),t.Hour(),t.Minute(),t.Second(),t.Nanosecond(),t.Location())
		break
	case time.September:
		date=time.Date(t.Year(),time.August,bigMonthConvert(t.Day()),t.Hour(),t.Minute(),t.Second(),t.Nanosecond(),t.Location())
		break
	case time.October:
		date=time.Date(t.Year(),time.September,bigMonthConvert(t.Day()),t.Hour(),t.Minute(),t.Second(),t.Nanosecond(),t.Location())

		break
	case time.November:
		date=time.Date(t.Year(),time.October,bigMonthConvert(t.Day()),t.Hour(),t.Minute(),t.Second(),t.Nanosecond(),t.Location())
		break
	case time.December:
		date=time.Date(t.Year(),time.November,bigMonthConvert(t.Day()),t.Hour(),t.Minute(),t.Second(),t.Nanosecond(),t.Location())
		break
	} 

	return &date
}

func bigMonthConvert(day int)int{
	if day==31{
		day=28
	}
	return day
}