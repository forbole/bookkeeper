package tables

import (
	"math"
	"math/big"

	"github.com/forbole/bookkeeper/coinApi"
	"github.com/forbole/bookkeeper/module/cosmos/utils"
	types "github.com/forbole/bookkeeper/types"
	tabletypes "github.com/forbole/bookkeeper/types/tabletypes"
)

func GetDateRewardCommissionValue(v tabletypes.RewardCommissions,denoms []types.Denom,vsCurrency string,lcd string)(
	[]tabletypes.DateRewardCommissionPrice,error){
	dateRewardCommissionPrice:=make([]tabletypes.DateRewardCommissionPrice,len(v))
	denomMap:=ConvertDenomToMap(denoms)
	for i,r:=range v{
		date,err:=utils.GetTimeByHeight(r.Height,lcd)
		if err!=nil{
			return nil,err
		}
		
		var price *big.Float
		if denomMap[r.Denom].Cointype=="crypto"{
			price,err=coinApi.GetCryptoPriceFromDate(*date,denomMap[r.Denom].CoinId,vsCurrency)
			if err!=nil{
				return nil,err
			}
		} else if denomMap[r.Denom].Cointype=="stablecoin"{
			price,err=coinApi.GetCurrencyPrice(denomMap[r.Denom].CoinId,vsCurrency)
			if err!=nil{
				return nil,err
			}
		}
		
		commission:=new(big.Float).Mul(new(big.Float).SetInt(r.Commission),denomMap[r.Denom].Exponent)
		reward:=new(big.Float).Mul(new(big.Float).SetInt(r.Reward),denomMap[r.Denom].Exponent)

		commissionPrice:=new(big.Float).Mul(commission,price)
		rewardPrice:=new(big.Float).Mul(reward,price)


		dateRewardCommissionPrice[i]=tabletypes.NewDateRewardCommissionPrice(*date,r.Reward,r.Commission,r.Denom,rewardPrice,commissionPrice)

	}
	return dateRewardCommissionPrice,nil
}

type denomDetails struct{
	Exponent *big.Float    `json:"exponent"`
    CoinId   string `json:"coin_id"`
    Cointype string `json:"cointype"`
}

func ConvertDenomToMap(denoms []types.Denom)map[string]denomDetails{
	denomDetailsMap:=make(map[string]denomDetails)
	for _,d:=range denoms{
		exponent := new(big.Float).SetFloat64((math.Pow(10, float64(-1*d.Exponent))))

		denomDetailsMap[d.Denom]=denomDetails{
			Exponent:exponent,
			CoinId  :d.CoinId,
			Cointype:d.Cointype,
		}
	}
	return denomDetailsMap
}