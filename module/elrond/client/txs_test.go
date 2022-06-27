package client_test

import (
	"testing"

	elrondclient "github.com/forbole/bookkeeper/module/elrond/client"
	"github.com/stretchr/testify/suite"
)

type ElrondClientTestSuite struct {
	suite.Suite

	client *elrondclient.ElrondClient
}

// We need this function to kick off the test suite, otherwise
// "go test" won't know about our tests
func TestElrondClientTestSuite(t *testing.T) {
	suite.Run(t, new(ElrondClientTestSuite))
}

func (suite *ElrondClientTestSuite) SetupTest() {
	suite.client = elrondclient.NewElrondClient("https://api.elrond.com")
}

func (suite *ElrondClientTestSuite) Test_GetSelfRedelegate() {

	address := "erd1q7mu5ek4nwgcszvzp9ycp9ehf6lay96dsap6j4luhdv9nt8m6smsmqmf3y"
	contract := "erd1qqqqqqqqqqqqqqqpqqqqqqqqqqqqqqqqqqqqqqqqqqqqq40llllsfjmn54"
	from := int64(1619564400)

	txs, err := suite.client.GetSelfRedelegate(address, contract, from)

	suite.Require().NoError(err)
	suite.Require().NotEmpty(txs)
}

func (suite *ElrondClientTestSuite) Test_GetTxs() {
	address := "erd1q7mu5ek4nwgcszvzp9ycp9ehf6lay96dsap6j4luhdv9nt8m6smsmqmf3y"

	txs, err := suite.client.GetTxs(address, "")
	suite.Require().NoError(err)
	suite.Require().NotEmpty(txs)
}

func (suite *ElrondClientTestSuite) Test_GetTxResult() {
	txHash := "cac935f6f3ebe0216d69a55ba08f1ad0ac96176404cbac1d5b26a8cde121ecb7"

	txResult, err := suite.client.GetTxResult(txHash)
	suite.Require().NoError(err)
	suite.Require().NotNil(txResult)
}
