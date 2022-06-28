package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	solanatypes "github.com/forbole/bookkeeper/module/solana/types"
	"github.com/rs/zerolog/log"
)

type SolanaBeachClient struct{
	api string
}

func NewSolanaBeachClient(api string)*SolanaBeachClient{
	return &(SolanaBeachClient{
		api: api,
	})
}

// GetStakeReward get stake eward from given address from now to given epoch
func (client *SolanaBeachClient) GetStakeReward(address string,epochFrom int)([]solanatypes.StakeReward,error){
	log.Trace().Str("module", "solana").Msg("GetStakeReward")

	var stakeReward []solanatypes.StakeReward
	size := 50
	for i := 0; ; i++ {

		query := fmt.Sprintf("v1/account/%s/stake-rewards?cursor=310",
		pubKey,size, i*size)

		bz, err := client.ping(query)
		if err != nil {
			return nil, err
		}

		var tx solanatypes.Stake
		err = json.Unmarshal(bz, &tx)
		if err != nil {
			return nil, fmt.Errorf("fail to marshal:%s", err)
		}
		stakeTxs = append(stakeTxs, tx.Data...)
		if tx.TotalPages==i+1{
			break
		}

	}

	return stakeTxs, nil
}

// GetStakeAccount get the staking accounts associate to the pubkey
func (client *SolanaBeachClient) GetStakeAccounts(pubKey string)([]solanatypes.StakeData,error){
	log.Trace().Str("module", "solana").Msg("GetStakes")

	var stakeTxs []solanatypes.StakeData
	size := 50
	for i := 0; ; i++ {

		query := fmt.Sprintf("v1/account/%s/stakes?limit=%d&offset=%d",
		pubKey,size, i*size)

		bz, err := client.ping(query)
		if err != nil {
			return nil, err
		}

		var tx solanatypes.Stake
		err = json.Unmarshal(bz, &tx)
		if err != nil {
			return nil, fmt.Errorf("fail to marshal:%s", err)
		}
		stakeTxs = append(stakeTxs, tx.Data...)
		if tx.TotalPages==i+1{
			break
		}

	}

	return stakeTxs, nil
}

func (client *SolanaBeachClient) ping(query string) ([]byte, error) {

	q := fmt.Sprintf("%s/%s", client.api, query)

	fmt.Println(q)
	var bz []byte
	resp, err := http.Get(q)
	if err != nil {
		return nil, fmt.Errorf("Fail to get tx from rpc:%s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Fail to get tx from rpc:Status :%s", resp.Status)
	}

	bz, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return bz, nil
}
