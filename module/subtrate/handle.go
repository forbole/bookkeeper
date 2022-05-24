package subtrate

import (
	"fmt"

	client "github.com/forbole/bookkeeper/module/subtrate/client"
	subtratetable "github.com/forbole/bookkeeper/module/subtrate/table"
	"github.com/forbole/bookkeeper/types"
)

func Handle(subtrate types.Subtrate) error {
	// create client
	client := client.NewSubscanClient(subtrate.ChainName)

	rewardSlash, err := subtratetable.GetRewardSlash(client, subtrate.Address[0])
	if err != nil {
		return err
	}
	fmt.Println(rewardSlash)
	return nil
}
