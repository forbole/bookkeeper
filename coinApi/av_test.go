package coinApi_test

import (
	"fmt"
	"os"

	"github.com/forbole/bookkeeper/coinApi"
)

func (suite *CoinApiTestSuite) Test_GetCurrencyPrice() {
	fmt.Printf("AV_API_KEY:%s", os.Getenv("AV_API_KEY"))
	price, err := coinApi.GetCurrencyPrice("hkd", "usd")
	suite.Require().NoError(err)
	suite.Require().NotZero(price)
}
