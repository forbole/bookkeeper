package table_test

import (
	"encoding/json"
	"testing"

	"github.com/forbole/bookkeeper/module/substrate/client"
	"github.com/forbole/bookkeeper/module/substrate/table"
	"github.com/forbole/bookkeeper/types"
	"github.com/stretchr/testify/suite"
)

type SubstrateTableTestSuite struct {
	suite.Suite

	client     *client.SubscanClient
	address    string
	denom      types.Denom
	vsCurrency string
	from       int64
}

// We need this function to kick off the test suite, otherwise
// "go test" won't know about our tests
func TestSubstrateTableTestSuite(t *testing.T) {
	suite.Run(t, new(SubstrateTableTestSuite))
}

/* 	client := client.NewSubscanClient(substrate.ChainName)

filename:=make([]string,len(substrate.Address)) */
func (suite *SubstrateTableTestSuite) SetupTest() {

	client := client.NewSubscanClient("polkadot")

	suite.client = client
	suite.address = "12L5PhJ2CT4MujSXoHTsBRZHQym4e6WYRhpAkgNWSwAnjZTf"
	denomStr := `{"denom":"DOT",
	"exponent":10,
	"coin_id":"polkadot",
	"cointype":"crypto"
  	}`

	var denom types.Denom
	err := json.Unmarshal([]byte(denomStr), &denom)
	suite.Require().NoError(err)
	suite.denom = denom
	suite.vsCurrency = "usd"
	suite.from = 1619564400
}

func (suite *SubstrateTableTestSuite) Test_GetCryptoPriceFromDate() {
	addressRewardPrice, err := table.GetRewardCommission(suite.client, suite.address, suite.denom,
		suite.vsCurrency, suite.from)
	suite.Require().NoError(err)
	suite.Require().NotEmpty(addressRewardPrice.Rows)
	suite.Require().Equal(addressRewardPrice.Address, suite.address)
}

func (suite *SubstrateTableTestSuite) Test_GetRewardSlash() {
	list, err := table.GetRewardSlash(suite.client, suite.address, suite.from)
	suite.Require().NoError(err)
	suite.Require().NotEmpty(list)
}
