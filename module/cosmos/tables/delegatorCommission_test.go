package tables_test

import (
	"math/big"

	"github.com/forbole/bookkeeper/module/cosmos/tables"
	"github.com/forbole/bookkeeper/types/tabletypes"
)

func (suite *CosmosTableTestSuite)Test_GetRewardCommission(){
	
	entries:=[]tabletypes.BalanceEntry{
		tabletypes.NewBalanceEntry(10357500,"9F828097E9E187091588F43BAFF9E2E353D19CC0DFBAF762C289FF4564FD4F10","10287229uatom","0","withdraw_delegator_reward"),
		tabletypes.NewBalanceEntry(10357500,"9F828097E9E187091588F43BAFF9E2E353D19CC0DFBAF762C289FF4564FD4F10","8414873uatom","0","withdraw_validator_commission"),
		tabletypes.NewBalanceEntry(10357500,"9F828097E9E187091588F43BAFF9E2E353D19CC0DFBAF762C289FF4564FD4F10","8414873uatom","0","otherMessage"),
	}

	address:="cosmosvaloper14kn0kk33szpwus9nh8n87fjel8djx0y070ymmj"
	
	addressEntries:=tabletypes.AddressBalanceEntry{
		Address: address,
		Rows: entries,
	}

	rewardCommissionExpected:=[]tabletypes.RewardCommission{
		tabletypes.NewRewardCommission("9F828097E9E187091588F43BAFF9E2E353D19CC0DFBAF762C289FF4564FD4F10",10357500,"uatom",big.NewInt(0),big.NewInt(10287229)),
		tabletypes.NewRewardCommission("9F828097E9E187091588F43BAFF9E2E353D19CC0DFBAF762C289FF4564FD4F10",10357500,"uatom",big.NewInt(8414873),big.NewInt(0)),

	}
	addressRewardCommissionExpected:=tabletypes.NewAddressRewardCommission(address,rewardCommissionExpected)
	
	rewardCommissionActual,err:=tables.GetRewardCommission(addressEntries)
	suite.Require().NoError(err)
	suite.Require().Equal(addressRewardCommissionExpected,*rewardCommissionActual)
}