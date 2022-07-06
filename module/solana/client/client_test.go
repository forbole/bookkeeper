package client_test

import (
	"testing"

	"github.com/forbole/bookkeeper/module/solana/client"

	"github.com/stretchr/testify/suite"
)

type SolanaClientTestSuite struct {
	suite.Suite

	solclient *(client.SolanaBeachClient)
}

func TestSolanaClientTestSuite(t *testing.T) {
	suite.Run(t, new(SolanaClientTestSuite))
}


func (suite *SolanaClientTestSuite) SetupTest() {
	c:=client.NewSolanaBeachClient("http://api.solanabeach.io")
	suite.solclient=c
}

func (suite *SolanaClientTestSuite) TestGetStakeReward() {
	address:="22i7vwvn9eNQU7FB7HqyyfLXUbVoh7yMBpMty61mDutM"
	epoch:=315
	_,err:=suite.solclient.GetStakeReward(address,epoch)
	suite.Require().NoError(err)
}

func (suite *SolanaClientTestSuite) TestGetEpochHistory() {
	_,err:=suite.solclient.GetEpochHistory()
	suite.Require().NoError(err)
}

func (suite *SolanaClientTestSuite) TestGetStakeAccounts() {
	pubkey:="DXRTh7JBgeaphmQVsdVKcafpWfznB12375MKEEDAEDLb"
	_,err:=suite.solclient.GetStakeAccounts(pubkey)
	suite.Require().NoError(err)
}
