package tables_test

import (
	"encoding/json"
	"testing"

	"github.com/forbole/bookkeeper/module/solana/client"
	"github.com/forbole/bookkeeper/module/solana/tables"

	"github.com/forbole/bookkeeper/types"
	"github.com/stretchr/testify/suite"
)

type SolanaTableTestSuite struct {
	suite.Suite
	testInput types.Solana
	period    types.Period
	solclient *(client.SolanaBeachClient)
}

func TestSolanaTableTestSuite(t *testing.T) {

	suite.Run(t, new(SolanaTableTestSuite))
}

func (suite *SolanaTableTestSuite) SetupTest() {
	chainStrings := `{
		"pubkey":"DXRTh7JBgeaphmQVsdVKcafpWfznB12375MKEEDAEDLb",
		"validator_identity":"forb5u56XgvzxiKfRt4FVNFQKJrd2LWAfNCsCqL6P7q",
		"solana_beach_api":"http://api.solanabeach.io",
		"denom":{"denom":"solana",
		  "exponent":9,
		  "coin_id":"solana",
		  "cointype":"crypto"
		}
		}`
	var chain types.Solana
	err := json.Unmarshal([]byte(chainStrings), &chain)
	if err != nil {
		panic(err)
	}

	c := client.NewSolanaBeachClient(chain.SolanaBeachApi)
	suite.solclient = c

	suite.testInput = chain
	suite.period = types.Period{
		From: 1619564400,
		To:   1651100400,
	}
}
func (suite *SolanaTableTestSuite) TestGetStakeRewardForPubKey() {
	vsCurrency := "usd"
	_, err := tables.GetStakeRewardForPubKey(suite.testInput, suite.period.From, vsCurrency, suite.solclient)
	suite.Require().NoError(err)
}
func (suite *SolanaTableTestSuite) TestGetRewardFromAddress() {
	address := "22i7vwvn9eNQU7FB7HqyyfLXUbVoh7yMBpMty61mDutM"
	epoch := 315
	vsCurrency := "usd"
	_, err := tables.GetRewardFromAddress(address, suite.testInput.Denom, epoch, vsCurrency, suite.solclient)
	suite.Require().NoError(err)
}
