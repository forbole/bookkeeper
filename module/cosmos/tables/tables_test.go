package tables_test

import (
	"encoding/json"
	"testing"

	"github.com/forbole/bookkeeper/types"
	"github.com/stretchr/testify/suite"
)

// We'll be able to store suite-wide
// variables and add methods to this
// test suite struct
type CosmosTableTestSuite struct {
	suite.Suite
	testInput types.IndividualChain
}

// We need this function to kick off the test suite, otherwise
// "go test" won't know about our tests
func TestCosmosTableTestSuite(t *testing.T) {
	suite.Run(t, new(CosmosTableTestSuite))
}

func (suite *CosmosTableTestSuite) SetupTest() {
	chainStrings := `{
      "chain_name":"cosmos",
      "denom":[{"denom":"uatom",
            "exponent":6,
            "coin_id":"cosmos",
            "cointype":"crypto"
            }],
      "grpc_endpoint":"adasdadsd",
      "rpc_endpoint":"https://rpc.cosmos.network",
      "lcd_endpoint":"https://api.cosmos.network",
      "validators":[
        {
        "validator_address":"123",
        "self_delegation_address":"abc"
        }],
      "fund_holding_account":["cosmosvaloper14kn0kk33szpwus9nh8n87fjel8djx0y070ymmj"]
    }`
	var chain types.IndividualChain
	json.Unmarshal([]byte(chainStrings), &chain)

	suite.testInput = chain
}
