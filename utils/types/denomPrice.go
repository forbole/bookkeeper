package types

import (
	"math/big"
)

type DenomPrice struct {
	Exponent *big.Float
	CoinId   string
	Price    *big.Float
}
