package utils_test

import (
	"testing"

	"github.com/forbole/bookkeeper/module/solana/client"
	"github.com/forbole/bookkeeper/module/solana/utils"

	"github.com/stretchr/testify/suite"
)

type SolanaUtilTestSuite struct {
	suite.Suite

	solclient *(client.SolanaBeachClient)
}

func TestSolanaUtilTestSuite(t *testing.T) {
	suite.Run(t, new(SolanaUtilTestSuite))
}

func (suite *SolanaUtilTestSuite) SetupTest() {
	c := client.NewSolanaBeachClient("http://api.solanabeach.io")
	suite.solclient = c
}

func (suite *SolanaUtilTestSuite) TestGetSelfDelegatorAddresses() {
	pubkey := "DXRTh7JBgeaphmQVsdVKcafpWfznB12375MKEEDAEDLb"
	validatorIdentity := "forb5u56XgvzxiKfRt4FVNFQKJrd2LWAfNCsCqL6P7q"
	strs, err := utils.GetSelfDelegatorAddresses(pubkey, validatorIdentity, suite.solclient)
	suite.Require().NoError(err)
	suite.Require().NotEmpty(strs)
}

func (suite *SolanaUtilTestSuite) TestGetEpochByTime() {
	from := int64(1619564400)
	_, err := utils.GetEpochByTime(from, suite.solclient)
	suite.Require().NoError(err)
}
