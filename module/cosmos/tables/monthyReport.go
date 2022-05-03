package cosmos

import (
	"fmt"
	"math/big"
	"time"

	"github.com/forbole/bookkeeper/module/cosmos/utils"
	types "github.com/forbole/bookkeeper/types"
	tabletypes "github.com/forbole/bookkeeper/types/tabletypes"
)

// GetMonthyReport get monthy report between certain period of time
func GetMonthyReport(details types.IndividualChain,period types.Period)([]tabletypes.AddressMonthyReport,error){
	var monthyReports []tabletypes.AddressMonthyReport

	from:=time.Unix(period.From,0)

	to:=time.Unix(period.To,0)

	balanceEntries,err:=GetTxs(details)
	if err!=nil{
		return nil,err
	}
	if balanceEntries==nil{
		return nil,nil
	}
	t:=to
	for _,b:=range balanceEntries{
		var monthyReportRows tabletypes.MonthyReportRows
		rewardCommission,err:=GetRewardCommission(b)
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
			rows:=rewardCommission.Rows
			recordForMonth :=make(map[string]*RewardCommission)
			for ;len(rows)>i && rows[i].Height>targetHeight;i++{
				// recordForMonth have denom as key and sum up reward and commission for the month 
				denomEntry,ok:=recordForMonth[rows[i].Denom]
				if !ok{
					e:=&RewardCommission{
						Commission:big.NewInt(0),
						Reward:big.NewInt(0),
					 }
					 denomEntry=e
				}

				commission:=big.NewInt(0).Add(denomEntry.Commission,rows[i].Commission)
				reward:=big.NewInt(0).Add(denomEntry.Reward,rows[i].Reward)
				denomEntry.Commission=new(big.Int).Set(commission)
				denomEntry.Reward=new(big.Int).Set(reward) 
				recordForMonth[rows[i].Denom]=denomEntry
				fmt.Println(denomEntry.Commission)
			}

			for key, element := range recordForMonth {
				to:=*(nextMonth(t))
				if time.Now().Before(*(nextMonth(t))){
					to=time.Now()
				}
				fmt.Println("Key:", key, "=>", "Element:", element)
				fmt.Println(t)
				fmt.Println(to)

				monthyReportRows=append(monthyReportRows,tabletypes.NewMonthyReportRow(t,to,element.Commission,element.Reward,key))
			}
			t=*(lastMonth(t))
	
		}
		monthyReports=append(monthyReports,tabletypes.NewAddressMonthyReport(b.Address,monthyReportRows))
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

type RewardCommission struct {
	Commission *big.Int
	Reward *big.Int
}