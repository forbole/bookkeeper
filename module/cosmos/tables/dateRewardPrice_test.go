package tables_test

import (
	"github.com/forbole/bookkeeper/module/cosmos/tables"
)

func (suite *CosmosTableTestSuite) Test_GetDateRewardValueFromDetails() {
	table, err := tables.GetDateRewardValueFromDetails(suite.testInput, suite.testPeriod, suite.testVsCurrency)
	suite.Require().NoError(err)
	suite.Require().NotEmpty(table)
}
