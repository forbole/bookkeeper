package utils

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/forbole/bookkeeper/module/solana/client"
)

// GetEpochByTime get the epoch near the timestamp
// if epoch is not found, it return the oldest epoch
func GetEpochByTime(time int64,client *client.SolanaBeachClient)(int,error){
	history,err:=client.GetEpochHistory()
	if err!=nil{
		return 0,err
	}
	
	// Declared an empty map interface
	var result map[string]string

	fmt.Println(string(history))
	// Unmarshal or Decode the JSON to the interface.
	err=json.Unmarshal(history, &result)
	if err!=nil{
		return 0,err
	}

	// find nesrest time (O(n))
	for i,r :=range result{
		fmt.Println(i)
		fmt.Println(r)

		t, err := strconv.Atoi(r)
		if err!=nil{
			return 0,err
		}

		if int64(t)>=time&&t!=0{
			epoch, err := strconv.Atoi(i)
			if err!=nil{
				return 0,err
			}
			return epoch,nil
		}

	}

	// Should not end up here
	return 0,fmt.Errorf("Cannot find epoch with time:%d",time)
	
}