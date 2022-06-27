package tables_test

import (
	"encoding/json"
	"testing"

	elrondclient "github.com/forbole/bookkeeper/module/elrond/client"
	"github.com/forbole/bookkeeper/module/elrond/tables"
	elrondtypes "github.com/forbole/bookkeeper/module/elrond/types"
	"github.com/forbole/bookkeeper/types"

	"github.com/stretchr/testify/suite"
)

type ElrondTableTestSuite struct {
	suite.Suite

	client *elrondclient.ElrondClient
}

// We need this function to kick off the test suite, otherwise
// "go test" won't know about our tests
func TestElrondTableTestSuite(t *testing.T) {
	suite.Run(t, new(ElrondTableTestSuite))
}

func (suite *ElrondTableTestSuite) SetupTest(){
	suite.client=elrondclient.NewElrondClient("https://api.elrond.com")
}

func (suite *ElrondTableTestSuite) Test_GetSelfRedelegate() {

	address:="erd1q7mu5ek4nwgcszvzp9ycp9ehf6lay96dsap6j4luhdv9nt8m6smsmqmf3y"
	contract:="erd1qqqqqqqqqqqqqqqpqqqqqqqqqqqqqqqqqqqqqqqqqqqqq40llllsfjmn54"
	from:=int64(1619564400)

	txs,err:=tables.GetTxs(suite.client,address,contract,from,"usd")

	suite.Require().NoError(err)
	suite.Require().NotEmpty(txs)
}

func (suite *ElrondTableTestSuite) Test_GetValueRow(){
	contract:="erd1qqqqqqqqqqqqqqqpqqqqqqqqqqqqqqqqqqqqqqqqqqqqq40llllsfjmn54"

	resultStr:=`{
		"txHash": "cac935f6f3ebe0216d69a55ba08f1ad0ac96176404cbac1d5b26a8cde121ecb7",
		"gasLimit": 12000000,
		"gasPrice": 1000000000,
		"gasUsed": 6075500,
		"miniBlockHash": "284f152c9caf846d4c0d5c0f73e3180ae5e2216186f625ff89511696aa0d1d9e",
		"nonce": 20,
		"receiver": "erd1qqqqqqqqqqqqqqqpqqqqqqqqqqqqqqqqqqqqqqqqqqqqq40llllsfjmn54",
		"receiverShard": 4294967295,
		"round": 9876552,
		"sender": "erd1q7mu5ek4nwgcszvzp9ycp9ehf6lay96dsap6j4luhdv9nt8m6smsmqmf3y",
		"senderShard": 1,
		"signature": "981eec2a24168e2cf563df6c6f53a58eedc5a1c4bda1cea6412062f4a1a9c21cea74da9634f9e24f92a2420f1954a49636d6e05fc257885f238c570c7f00e408",
		"status": "success",
		"value": "0",
		"fee": "135500000000000",
		"timestamp": 1655376912,
		"data": "cmVEZWxlZ2F0ZVJld2FyZHM=",
		"function": "reDelegateRewards",
		"action": {
		  "category": "stake",
		  "name": "reDelegateRewards",
		  "description": "Redelegate rewards from staking provider Forbole",
		  "arguments": {
			"providerName": "Forbole",
			"providerAvatar": "https://s3.amazonaws.com/keybase_processed_uploads/f5b0771af36b2e3d6a196a29751e1f05_360_360.jpeg"
		  }
		},
		"results": [
		  {
			"hash": "0ccd3253b2dd76c6b297df21301dbb8cd2d15f3b9fcf008378514edfec21d6cf",
			"timestamp": 1655376912,
			"nonce": 0,
			"gasLimit": 0,
			"gasPrice": 1000000000,
			"value": "340845645235676530768",
			"sender": "erd1qqqqqqqqqqqqqqqpqqqqqqqqqqqqqqqqqqqqqqqqqqqqq40llllsfjmn54",
			"receiver": "erd1qqqqqqqqqqqqqqqpqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqplllst77y4l",
			"prevTxHash": "cac935f6f3ebe0216d69a55ba08f1ad0ac96176404cbac1d5b26a8cde121ecb7",
			"originalTxHash": "cac935f6f3ebe0216d69a55ba08f1ad0ac96176404cbac1d5b26a8cde121ecb7",
			"callType": "0",
			"miniBlockHash": "4e9924e514c1bc15afb8b033d0ebb045dd08a0e97333924a38f275a44031b5bd"
		  },
		  {
			"hash": "65183e62a48e16a2150129fcc2a13f00f2498e87a2754449400d96ea4d6401af",
			"timestamp": 1655376924,
			"nonce": 21,
			"gasLimit": 0,
			"gasPrice": 1000000000,
			"value": "59245000000000",
			"sender": "erd1qqqqqqqqqqqqqqqpqqqqqqqqqqqqqqqqqqqqqqqqqqqqq40llllsfjmn54",
			"receiver": "erd1q7mu5ek4nwgcszvzp9ycp9ehf6lay96dsap6j4luhdv9nt8m6smsmqmf3y",
			"data": "QDZmNmI=",
			"prevTxHash": "cac935f6f3ebe0216d69a55ba08f1ad0ac96176404cbac1d5b26a8cde121ecb7",
			"originalTxHash": "cac935f6f3ebe0216d69a55ba08f1ad0ac96176404cbac1d5b26a8cde121ecb7",
			"callType": "0",
			"miniBlockHash": "d3be0615ab77ad048f1f4f1f9991dfcf250a79e50c778faee6e4e6d40ec6bcb4",
			"logs": {
			  "address": "erd1q7mu5ek4nwgcszvzp9ycp9ehf6lay96dsap6j4luhdv9nt8m6smsmqmf3y",
			  "events": [
				{
				  "address": "erd1q7mu5ek4nwgcszvzp9ycp9ehf6lay96dsap6j4luhdv9nt8m6smsmqmf3y",
				  "identifier": "completedTxEvent",
				  "topics": [
					"ysk19vPr4CFtaaVboI8a0KyWF2QEy6wdWyaozeEh7Lc="
				  ],
				  "order": 0
				}
			  ]
			}
		  }
		],
		"price": 50.37,
		"logs": {
		  "address": "erd1qqqqqqqqqqqqqqqpqqqqqqqqqqqqqqqqqqqqqqqqqqqqq40llllsfjmn54",
		  "events": [
			{
			  "address": "erd1q7mu5ek4nwgcszvzp9ycp9ehf6lay96dsap6j4luhdv9nt8m6smsmqmf3y",
			  "identifier": "delegate",
			  "topics": [
				"Enou/HbXUEhQ",
				"kdjMCEEBb7lr",
				"Avk=",
				"DXc0vfrWHe3nsQ=="
			  ],
			  "order": 0
			}
		  ]
		},
		"operations": [
		  {
			"id": "0ccd3253b2dd76c6b297df21301dbb8cd2d15f3b9fcf008378514edfec21d6cf",
			"action": "transfer",
			"type": "egld",
			"sender": "erd1qqqqqqqqqqqqqqqpqqqqqqqqqqqqqqqqqqqqqqqqqqqqq40llllsfjmn54",
			"receiver": "erd1qqqqqqqqqqqqqqqpqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqplllst77y4l",
			"value": "340845645235676530768"
		  }
		]
	  }`

	var expectedTx elrondtypes.TxResult
	err:=json.Unmarshal([]byte(resultStr),&expectedTx)
	suite.Require().NoError(err)

	input:=[]elrondtypes.TxResult{
		expectedTx,
	}

	denomStr:=`{"denom":"egld",
      "exponent":18,
      "coin_id":"elrond-erd-2",
      "cointype":"crypto"
    }`

	var denom types.Denom
	err=json.Unmarshal([]byte(denomStr),&denom)
	suite.Require().NoError(err)

	
	rows,err:=tables.GetValueRow(input,"usd",denom,contract)
	suite.Require().NoError(err)
	suite.Require().NotEmpty(rows)
}

