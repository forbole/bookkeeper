package client_test

import (
	"testing"

	"github.com/forbole/bookkeeper/module/substrate/client"
	"github.com/stretchr/testify/suite"
)

type SubstrateClientTestSuite struct {
	suite.Suite

	api string
}

// We need this function to kick off the test suite, otherwise
// "go test" won't know about our tests
func TestSubstrateTableTestSuite(t *testing.T) {
	suite.Run(t, new(SubstrateClientTestSuite))
}

/* 	client := client.NewSubscanClient(substrate.ChainName)

filename:=make([]string,len(substrate.Address)) */
func (suite *SubstrateClientTestSuite) SetupTest() {
	suite.api = "polkadot"
}

func (suite *SubstrateClientTestSuite) NewSubscanClient() {
	c := client.NewSubscanClient(suite.api)
	suite.Require().NotNil(c)
}

func (suite *SubstrateClientTestSuite) Test_GetRewardSlash() {
	c := client.NewSubscanClient(suite.api)
	requestUrl := "/api/now"

	var timestamp timestamp
	err := c.CallApi(requestUrl, nil, &timestamp)
	suite.Require().NoError(err)
	suite.Require().NotEmpty(timestamp)
}

type timestamp struct {
	Code        int    `json:"code"`
	Message     string `json:"message"`
	GeneratedAt int    `json:"generated_at"`
	Data        int    `json:"data"`
}

func (v *timestamp) SubscanApi() {}
