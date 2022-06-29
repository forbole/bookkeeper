package utils

import (
	"encoding/json"
	"fmt"

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
	var result map[string]int

	// Unmarshal or Decode the JSON to the interface.
	json.Unmarshal(history, &result)

	// find nesrest time (O(n))
	for _,r :=range result{
		if r>int(time) && r!=0{
			return r,nil
		}
	}

	// Should not end up here
	return 0,fmt.Errorf("Cannot find epoch with time:%d",time)
	
}