package utils

import (
	"math"
	"math/big"

	"github.com/forbole/bookkeeper/types"
	utilstypes "github.com/forbole/bookkeeper/utils/types"

	"github.com/forbole/bookkeeper/coinApi"
)

// ConvertAttributeToMap turn attribute into a map so that it is easy to find attributes
// Make Denom as a key and query their price
func ConvertDenomToMap(denom []types.Denom,vsCurrency string)(map[string]*utilstypes.DenomPrice,error){
	coinPrice:=make(map[string]*utilstypes.DenomPrice)

	for _,d:=range denom{
		exponent:=new(big.Float).SetFloat64((math.Pow(10,float64(-1 * d.Exponent))))

		if d.Cointype=="stablecoin"{
			price,err:=coinApi.GetPriceFromAV(d.CoinId,vsCurrency)
			if err!=nil{
				return nil,err
			}
			coinPrice[d.Denom]=&utilstypes.DenomPrice{
				Exponent: exponent,
				CoinId: d.CoinId,
				Price: price,
			}
		}else if d.Cointype=="crypto"{
			price,err:=coinApi.GetPriceFromCoingecko(d.CoinId,vsCurrency)
			if err!=nil{
				return nil,err
			}
			coinPrice[d.Denom]=&utilstypes.DenomPrice{
				Exponent: exponent,
				CoinId: d.CoinId,
				Price: price,
			}
		}
	}
	return coinPrice,nil
}