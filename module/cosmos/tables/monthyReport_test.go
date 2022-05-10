package tables_test

import (
	"github.com/forbole/bookkeeper/module/cosmos/tables"
	"github.com/forbole/bookkeeper/types"
)

func (suite *CosmosTableTestSuite) Test_MonthyReport() {
	period := types.Period{
		From: 1619564400,
		To:   1651100400,
	}
	reportActual, err := tables.GetMonthyReport(suite.testInput, period)
	suite.Require().NoError(err)
	suite.Require().NotEmpty(reportActual)
}
