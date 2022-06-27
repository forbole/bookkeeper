package subtrate

import (
	"fmt"
	"io/ioutil"

	client "github.com/forbole/bookkeeper/module/subtrate/client"
	subtratetable "github.com/forbole/bookkeeper/module/subtrate/table"
	"github.com/forbole/bookkeeper/types"
)

func Handle(subtrate types.Subtrate, vsCurrency string, outputFolder string, period types.Period) ([]string, error) {
	// create client
	client := client.NewSubscanClient(subtrate.ChainName)

	filename := make([]string, len(subtrate.Address))

	for i, address := range subtrate.Address {
		rewardSlash, err := subtratetable.GetRewardCommission(client, address, subtrate.Denom[0], vsCurrency, period.From)
		if err != nil {
			return nil, err
		}

		outputcsv2 := rewardSlash.GetCSV()

		filename2 := fmt.Sprintf("%s/%s_%s_reward_price.csv", outputFolder, subtrate.ChainName, address)
		err = ioutil.WriteFile(filename2, []byte(outputcsv2), 0600)
		if err != nil {
			return nil, err
		}

		filename[i] = filename2
	}

	return filename, nil
}
