package coinApi_test

import (
	"os"

	"github.com/forbole/bookkeeper/coinApi"
)

func (suite *CoinApiTestSuite) Test_GetCurrencyPrice(){
	os.Setenv("AV_API_KEY","RIJ86K6A7GFAWR4N")
	price,err:=coinApi.GetCurrencyPrice("hkd","usd")
	suite.Require().NoError(err)
	suite.Require().NotZero(price)
}

