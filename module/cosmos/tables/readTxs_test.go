package tables_test

import (
	"encoding/json"

	"github.com/forbole/bookkeeper/module/cosmos/tables"
	"github.com/forbole/bookkeeper/types"
)

// This is an example test that will always succeed
func (suite *CosmosTableTestSuite) ReadTx_Test() {
	chainStrings:=`{
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

  	from:=int64(1619564400)

 	var chain types.IndividualChain
	json.Unmarshal([]byte(chainStrings),&chain)

	balanceEntries,err:=tables.GetTxs(chain,from)
	suite.Require().NoError(err)
	suite.Require().True(len(balanceEntries)>0)
}