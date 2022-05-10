package utils_test

import (
	"testing"
	"time"

	"github.com/forbole/bookkeeper/module/cosmos/utils"
	"github.com/stretchr/testify/suite"
)

// We'll be able to store suite-wide
// variables and add methods to this
// test suite struct
type CosmosUtilsTestSuite struct {
	suite.Suite
}

// We need this function to kick off the test suite, otherwise
// "go test" won't know about our tests
func TestCosmosUtilsTestSuite(t *testing.T) {
	suite.Run(t, new(CosmosUtilsTestSuite))
}

func (suite *CosmosUtilsTestSuite) Test_GetHeightByDate() {
	targetTime := time.Unix(1619564400, 0)
	lcd := "https://api.cosmos.network"
	_, err := utils.GetHeightByDate(targetTime, lcd)
	suite.Assert().NoError(err)
}

func (suite *CosmosUtilsTestSuite) Test_GetDateByHeight() {
	height := 10361414
	lcd := "https://api.cosmos.network"
	t, err := utils.GetTimeByHeight(height, lcd)
	suite.Assert().NoError(err)

	expectedTime, err := time.Parse(time.RFC3339, "2022-05-04T11:40:01.438244256Z")
	suite.Assert().NoError(err)

	suite.Assert().True(expectedTime.Equal(*t))
}
