package utils_test

import (
	"testing"
	"time"

	"github.com/forbole/bookkeeper/module/substrate/client"
	"github.com/forbole/bookkeeper/module/substrate/utils"
	"github.com/stretchr/testify/suite"
)

type SubstrateUtilTestSuite struct {
	suite.Suite

	client *client.SubscanClient
}

// We need this function to kick off the test suite, otherwise
// "go test" won't know about our tests
func TestSubstrateUtilTestSuite(t *testing.T) {
	suite.Run(t, new(SubstrateUtilTestSuite))
}

/* 	client := client.NewSubscanClient(substrate.ChainName)

filename:=make([]string,len(substrate.Address)) */
func (suite *SubstrateUtilTestSuite) SetupTest() {

	client := client.NewSubscanClient("polkadot")

	suite.client = client

}

func (suite *SubstrateUtilTestSuite) Test_GetTimeByBlockNum() {
	//2022-06-17 12:42:00 (+UTC)
	expectedTimestamp := time.Date(2022, 6, 17, 12, 42, 00, 00, time.UTC)
	timestamp, err := utils.GetTimeByBlockNum(10781140, suite.client)
	suite.Require().NoError(err)
	suite.Require().True(expectedTimestamp.Equal(*timestamp))
}
