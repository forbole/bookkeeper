package tables_test

import (
	"fmt"
	"io/ioutil"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
	"time"

	database "github.com/forbole/bookkeeper/module/flow/db"
	"github.com/forbole/bookkeeper/module/flow/tables"
	client "github.com/forbole/bookkeeper/module/flow/utils"

	"github.com/forbole/bookkeeper/types"
	"github.com/stretchr/testify/suite"
)

func TestFlowRewardCommissionTestSuite(t *testing.T) {
	suite.Run(t, new(FlowRewardCommissionTestSuite))
}

type FlowRewardCommissionTestSuite struct {
	suite.Suite

	database *database.FlowDb
}

func (suite *FlowRewardCommissionTestSuite) SetupTest() {
	// Build the database
	cfg := types.NewDatabase(
		"localhost",
		5433,
		"bdjuno",
		"bdjuno", "disable",
		"public",
		"password",
	)

	db, err := database.Build(cfg)
	suite.Require().NoError(err)

	// Delete the public schema
	_, err = db.Sql.Exec(`DROP SCHEMA public CASCADE;`)
	suite.Require().NoError(err)

	// Re-create the schema
	_, err = db.Sql.Exec(`CREATE SCHEMA public;`)
	suite.Require().NoError(err)

	dirPath := path.Join("..", "db", "schema")
	dir, err := ioutil.ReadDir(dirPath)
	suite.Require().NoError(err)

	for _, fileInfo := range dir {
		file, err := ioutil.ReadFile(filepath.Join(dirPath, fileInfo.Name()))
		suite.Require().NoError(err)

		commentsRegExp := regexp.MustCompile(`/\*.*\*/`)
		requests := strings.Split(string(file), ";")
		for _, request := range requests {
			_, err := db.Sql.Exec(commentsRegExp.ReplaceAllString(request, ""))
			suite.Require().NoError(err)
		}
	}

	suite.database = db
}

func (suite *FlowRewardCommissionTestSuite) Test_GetRewardCommission() {

	//17 June 2022, 15:34:50
	timestamp := time.Date(2022, 6, 17, 15, 34, 50, 0, time.UTC)
	payer := "645007cf9b780ffd"
	amount := 10.01
	eventType := "A.8624b52f9ddcd04a.FlowIDTableStaking.RewardTokensWithdrawn"
	height := int64(31836616)
	transactionId := "d502d00cb4f9b4ba5c9479dbae3dc5dd10de68523077cd7b21d509f82cab7378"
	eventValue := fmt.Sprintf("%s(payer: %s, amount:%f)", eventType, payer, amount)
	jsonPath := []byte("{}")

	// set up db
	_, err := suite.database.Sql.Exec(`
	INSERT INTO block(height,id,parent_id,collection_guarantees,timestamp) VALUES 
	($1,$2,$3,$4,$5)`, height, "", "", jsonPath, timestamp)
	suite.Require().NoError(err)

	_, err = suite.database.Sql.Exec(`
	INSERT INTO transaction 
		(height,transaction_id,script,arguments,reference_block_id,gas_limit,proposal_key ,payer,authorizers,payload_signature,envelope_signatures) 
	VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)`,
		height, transactionId, "", nil,
		"",
		9999, "", payer, nil, jsonPath, jsonPath)
	suite.Require().NoError(err)

	_, err = suite.database.Sql.Exec(`
	INSERT INTO event(height,type,transaction_id,transaction_index,event_index,value)
	VALUES ($1,$2,$3,$4,$5,$6)`, 1, eventType, transactionId, "", 1, eventValue)
	suite.Require().NoError(err)

	// flow client
	flowEndpoint := "access.mainnet.nodes.onflow.org:9000"
	lastSpork := 16
	vsCurrency := "usd"

	flowclient, err := client.NewFlowClient(flowEndpoint, lastSpork)
	suite.Require().NoError(err)

	val, err := tables.GetRewardCommission(payer, suite.database, flowclient, vsCurrency)
	suite.Require().NoError(err)

	suite.Require().Equal(payer, val.Address)
	suite.Require().NotEmpty(val.Rows)
}
