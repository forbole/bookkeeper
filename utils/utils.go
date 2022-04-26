package utils

import (
	utilstypes "github.com/forbole/bookkeeper/utils/types"
	"github.com/forbole/bookkeeper/types"

	"github.com/forbole/bookkeeper/coinApi"

)

// ConvertAttributeToMap turn attribute into a map so that it is easy to find attributes
// Make Denom as a key and query their price
func ConvertDenomToMap(denom []types.Denom,vsCurrency string)(map[string]*utilstypes.DenomPrice,error){
	coinPrice:=make(map[string]*utilstypes.DenomPrice)

	for _,d:=range denom{
		if d.Stablecoin{
			price,err:=coinApi.GetPriceFromAV(d.CoinId,vsCurrency)
			if err!=nil{
				return nil,err
			}
			coinPrice[d.Denom]=&utilstypes.DenomPrice{
				Denom: d,
				Price: price,
			}
		}else{
			price,err:=coinApi.GetPriceFromCoingecko(d.CoinId,vsCurrency)
			if err!=nil{
				return nil,err
			}
			coinPrice[d.Denom]=&utilstypes.DenomPrice{
				Denom: d,
				Price: price,
			}
		}
	}
	return coinPrice,nil
}