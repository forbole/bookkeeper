package tables_test

import (
	"encoding/json"
	"testing"

	"github.com/forbole/bookkeeper/module/flow/tables"
	"github.com/forbole/bookkeeper/types"
	"github.com/stretchr/testify/suite"
)

// We'll be able to store suite-wide
// variables and add methods to this
// test suite struct
type FlowTableTestSuite struct {
	suite.Suite
	testInput types.Flow
	period    types.Period
}

// We need this function to kick off the test suite, otherwise
// "go test" won't know about our tests
func TestFlowTableTestSuite(t *testing.T) {

	suite.Run(t, new(FlowTableTestSuite))
}

func (suite *FlowTableTestSuite) SetupTest() {
	chainStrings := `{
		"flowjuno":"https://gql.flow.forbole.com/v1/graphql",
		"flow_endpoint":"access.mainnet.nodes.onflow.org:9000",
		"addresses":["fb397444147918de"],
		"denom":"flow",
		"exponent":8,
		"last_spork":16
	  }`
	var chain types.Flow
	err := json.Unmarshal([]byte(chainStrings), &chain)
	if err != nil {
		panic(err)
	}

	suite.testInput = chain
	suite.period = types.Period{
		From: 1619564400,
		To:   1651100400,
	}
}

func (suite *FlowTableTestSuite) Test_GetInfoFromAddress() {
	startHeight := uint64(21291692)
	nodeInfo, err := tables.GetNodeInfoFromAddress(suite.testInput.Addresses[0],
		suite.testInput.FlowJuno, startHeight)
	suite.Require().NoError(err)
	suite.Require().NotEmpty(nodeInfo)
}
