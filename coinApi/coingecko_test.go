package coinApi_test

import (
	"math/big"
	"testing"
	"time"

	"github.com/forbole/bookkeeper/coinApi"
	"github.com/stretchr/testify/suite"
)

type CoinApiTestSuite struct {
	suite.Suite
}

// We need this function to kick off the test suite, otherwise
// "go test" won't know about our tests
func TestCoinApiTestSuite(t *testing.T) {
	suite.Run(t, new(CoinApiTestSuite))
}

func (suite *CoinApiTestSuite) Test_GetCryptoPriceFromDate() {
	date := time.Date(2017, 1, 30, 0, 0, 0, 0, time.UTC)

	expectedPrice := new(big.Float).SetFloat64(920.9911458822621)

	price, err := coinApi.GetCryptoPriceFromDate(date, "bitcoin", "usd")
	suite.Require().NoError(err)
	suite.Require().Equal(expectedPrice, price)

}

func (suite *CoinApiTestSuite) Test_GetCryptoPrice() {
	price, err := coinApi.GetCryptoPrice("bitcoin", "usd")
	suite.Require().NoError(err)
	suite.Require().NotZero(price)
}
