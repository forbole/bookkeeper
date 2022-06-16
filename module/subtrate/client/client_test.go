package client_test

import (
	"testing"

	"github.com/forbole/bookkeeper/module/subtrate/client"
	"github.com/forbole/bookkeeper/module/subtrate/types"
	"github.com/stretchr/testify/suite"
)

type SubtrateClientTestSuite struct {
	suite.Suite

	api string
}

// We need this function to kick off the test suite, otherwise
// "go test" won't know about our tests
func TestSubtrateTableTestSuite(t *testing.T) {
	suite.Run(t, new(SubtrateClientTestSuite))
}
/* 	client := client.NewSubscanClient(subtrate.ChainName)

	filename:=make([]string,len(subtrate.Address)) */
func (suite *SubtrateClientTestSuite) SetupTest(){
	suite.api="polkadot"
}

func (suite *SubtrateClientTestSuite) NewSubscanClient() {
	c:=client.NewSubscanClient(suite.api)
	suite.Require().NotNil(c)
}

func (suite *SubtrateClientTestSuite) Test_GetRewardSlash() {
	c:=client.NewSubscanClient(suite.api)
	requestUrl := "/api/scan/account/reward_slash"

	type Payload struct {
		Row     int    `json:"row"`
		Page    int    `json:"page"`
		Address string `json:"address"`
	}

	payload:=Payload{
		Row: 1,
		Page: 1,
		Address: "12L5PhJ2CT4MujSXoHTsBRZHQym4e6WYRhpAkgNWSwAnjZTf",
	}
	
	var rewardSlash types.RewardSlash
	err:=c.CallApi(requestUrl,payload,&rewardSlash)
	suite.Require().NoError(err)
	suite.Require().NotEmpty(rewardSlash)
}