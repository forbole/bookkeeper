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
	testInput      types.CosmosDetails
	testVsCurrency string
	testPeriod     types.Period
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
        "validator_address":"cosmosvaloper14kn0kk33szpwus9nh8n87fjel8djx0y070ymmj",
        "self_delegation_address":"abc"
        }],
      "fund_holding_account":["cosmos1kvp570cd6zvzh8ffrhz7lmytt6v6u2gx393tla"]
    }`
	var chain types.CosmosDetails
	err := json.Unmarshal([]byte(chainStrings), &chain)
	if err != nil {
		panic(err)
	}

	suite.testInput = chain
	/*
	   "period":{"from":1619564400,
	   "to":1651100400}
	*/
	suite.testPeriod = types.Period{
		From: 1619564400,
		To:   1651100400,
	}

	suite.testVsCurrency = "USD"
}
