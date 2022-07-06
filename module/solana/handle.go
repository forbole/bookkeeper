package solana

import (
	"fmt"
	"io/ioutil"

	"github.com/forbole/bookkeeper/module/solana/client"
	"github.com/forbole/bookkeeper/module/solana/tables"
	"github.com/forbole/bookkeeper/types"
)

func HandleReward(solana types.Solana, period types.Period, vsCurrency string, outputFile string) ([]string, error) {
	solClient := client.NewSolanaBeachClient(solana.SolanaBeachApi)
	addressRewardPriceTables, err := tables.GetStakeRewardForPubKey(solana, period.From, vsCurrency, solClient)
	if err != nil {
		return nil, err
	}
	var filenames []string
	for _, table := range addressRewardPriceTables {
		csv := table.GetCSV()
		filename := fmt.Sprintf("%s/solana_%s_reward.csv", outputFile, table.Address)
		err = ioutil.WriteFile(filename, []byte(csv), 0600)
		if err != nil {
			return nil, err
		}
		filenames = append(filenames, filename)
	}
	return filenames, nil
}
