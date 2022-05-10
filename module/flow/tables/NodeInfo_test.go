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
		"node_ids":["fb397444147918de"],
		"denom":"flow",
		"exponent":8,
		"last_spork":16
	  }`
	var chain types.Flow
	json.Unmarshal([]byte(chainStrings), &chain)

	suite.testInput = chain
}

func (suite *FlowTableTestSuite) Test_GetNodeInfo() {
	nodeInfo, err := tables.GetNodeInfo(suite.testInput.Addresses[0], suite.testInput.FlowJuno)
	suite.Require().NoError(err)
	suite.Require().NotEmpty(nodeInfo)
}
