package subtrate

import (
	"fmt"
	"io/ioutil"

	client "github.com/forbole/bookkeeper/module/subtrate/client"
	subtratetable "github.com/forbole/bookkeeper/module/subtrate/table"
	"github.com/forbole/bookkeeper/types"
)

func Handle(subtrate types.Subtrate,vsCurrency string,outputFolder string,period types.Period) (string,error) {
	// create client
	client := client.NewSubscanClient(subtrate.ChainName)

	rewardSlash, err := subtratetable.GetRewardCommission(client, subtrate.Address[0],subtrate.Denom[0],vsCurrency,period.From)
	if err != nil {
		return "", err
	}
	fmt.Println(rewardSlash.GetCSV())

	outputcsv2 := rewardSlash.GetCSV()

	filename2 := fmt.Sprintf("%s/%s_%s_reward_price.csv", outputFolder, subtrate.ChainName, subtrate.Address[0])
	err = ioutil.WriteFile(filename2, []byte(outputcsv2), 0600)
	if err != nil {
		return "", err
	}

	return filename2,nil
}
