package substrate

import (
	"fmt"
	"io/ioutil"

	client "github.com/forbole/bookkeeper/module/substrate/client"
	substratetable "github.com/forbole/bookkeeper/module/substrate/table"
	"github.com/forbole/bookkeeper/types"
)

func Handle(substrate types.Substrate, vsCurrency string, outputFolder string, period types.Period) ([]string, error) {
	// create client
	client := client.NewSubscanClient(substrate.ChainName)

	filename := make([]string, len(substrate.Address))

	for i, address := range substrate.Address {
		rewardSlash, err := substratetable.GetRewardCommission(client, address, substrate.Denom[0], vsCurrency, period.From)
		if err != nil {
			return nil, err
		}

		outputcsv2 := rewardSlash.GetCSV()

		filename2 := fmt.Sprintf("%s/%s_%s_reward_price.csv", outputFolder, substrate.ChainName, address)
		err = ioutil.WriteFile(filename2, []byte(outputcsv2), 0600)
		if err != nil {
			return nil, err
		}

		filename[i] = filename2
	}

	return filename, nil
}
