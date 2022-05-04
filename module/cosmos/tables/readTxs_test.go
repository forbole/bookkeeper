package tables_test

import (
	"github.com/forbole/bookkeeper/module/cosmos/tables"
)

// This is an example test that will always succeed
func (suite *CosmosTableTestSuite) Test_ReadTxs() {
	from:=int64(1619564400)
	balanceEntries,err:=tables.GetTxs(suite.testInput,from)
	suite.Require().NoError(err)
	suite.Require().True(len(balanceEntries)>0)
}