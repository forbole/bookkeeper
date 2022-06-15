package db_test

import (
	"io/ioutil"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
	"time"

	database "github.com/forbole/bookkeeper/module/flow/db"
	types "github.com/forbole/bookkeeper/types"
	"github.com/stretchr/testify/suite"
)

func TestDatabaseTestSuite(t *testing.T) {
	suite.Run(t, new(DbTestSuite))
}

type DbTestSuite struct {
	suite.Suite

	database *database.FlowDb
}

func (suite *DbTestSuite) SetupTest() {
	// Build the database
	cfg := types.NewDatabase(
				"localhost",
				5433,
				"flowjuno",
				"","disable",
				"public",
				"",
	)

	db, err := database.Build(cfg)
	suite.Require().NoError(err)

	// Delete the public schema
	_, err = db.Sql.Exec(`DROP SCHEMA public CASCADE;`)
	suite.Require().NoError(err)

	// Re-create the schema
	_, err = db.Sql.Exec(`CREATE SCHEMA public;`)
	suite.Require().NoError(err)

	dirPath := path.Join(".", "schema")
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

func (suite *DbTestSuite) Test_GetWithdrawReward(){
	_,err:=suite.database.Sql.Exec(`
	INSERT INTO block(height,id,parent_id,collection_guarantees,timestamp) VALUES 
	$1,$2,$3,$4,$5`,1,"","","",time.Date(2020,1,1,1,1,1,1,time.UTC))
	suite.Require().NoError(err)

	_, err = suite.database.Sql.Exec(`
	INSERT INTO transaction 
		(height,transaction_id,script,arguments,reference_block_id,gas_limit,proposal_key ,payer,authorizers,payload_signature,envelope_signatures) 
	VALUES $1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11`,1,"d502d00cb4f9b4ba5c9479dbae3dc5dd10de68523077cd7b21d509f82cab7378","",nil,
	"852bc4a6e27c326dc1d270677a25340f2b7990fad3da29e2ce9f9c2c5c67d9e2",
	9999,"","645007cf9b780ffd",nil,"","")
	suite.Require().NoError(err)


}
