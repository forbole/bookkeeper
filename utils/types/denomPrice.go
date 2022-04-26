package types

import (
	"math/big"

	"github.com/forbole/bookkeeper/types"
)

type DenomPrice struct{
	types.Denom
	Price *big.Float
}
