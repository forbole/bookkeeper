package client

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"os"
	"time"

	solanatypes "github.com/forbole/bookkeeper/module/solana/types"
	"github.com/rs/zerolog/log"
)

type SolanaBeachClient struct {
	api string
}

func NewSolanaBeachClient(api string) *SolanaBeachClient {
	return &(SolanaBeachClient{
		api: api,
	})
}

func (client *SolanaBeachClient) GetEpochHistory() ([]byte, error) {
	log.Trace().Str("module", "solana").Msg("GetEpochHistory")

	query := "v1/epoch-history"
	return client.get(query)
}

// GetStakeReward get stake eward from given address from now to given epoch
func (client *SolanaBeachClient) GetStakeReward(address string, epochFrom int) ([]solanatypes.StakeReward, error) {
	log.Trace().Str("module", "solana").Msg("GetStakeReward")

	cursor := ""
	var stakeReward []solanatypes.StakeReward
	currentEpoch := math.MaxInt
	for currentEpoch >= epochFrom {

		query := fmt.Sprintf("v1/account/%s/stake-rewards%s",
			address, cursor)

		bz, err := client.get(query)
		if err != nil {
			return nil, err
		}

		var tx []solanatypes.StakeReward
		err = json.Unmarshal(bz, &tx)
		if err != nil {
			return nil, fmt.Errorf("fail to marshal:%s", err)
		}
		if len(tx) == 0 {
			return stakeReward, nil
		}
		stakeReward = append(stakeReward, tx...)

		currentEpoch := tx[len(tx)-1].Epoch
		cursor = fmt.Sprintf("?cursor=%d", currentEpoch)
	}

	return stakeReward, nil
}

// GetStakeAccount get the staking accounts associate to the pubkey
func (client *SolanaBeachClient) GetStakeAccounts(pubKey string) ([]solanatypes.StakeData, error) {
	log.Trace().Str("module", "solana").Msg("GetStakeAccounts")

	var stakeTxs []solanatypes.StakeData
	size := 50
	for i := 0; ; i++ {

		query := fmt.Sprintf("v1/account/%s/stakes?limit=%d&offset=%d",
			pubKey, size, i*size)

		bz, err := client.get(query)
		if err != nil {
			return nil, err
		}

		var tx solanatypes.Stake
		err = json.Unmarshal(bz, &tx)
		if err != nil {
			fmt.Println(string(bz))
			return nil, fmt.Errorf("fail to marshal:%s", err)
		}
		stakeTxs = append(stakeTxs, tx.Data...)
		if tx.TotalPages == i+1 {
			break
		}

	}

	return stakeTxs, nil
}

func (client *SolanaBeachClient) get(query string) ([]byte, error) {

	q := fmt.Sprintf("%s/%s", client.api, query)
	fmt.Println(q)
	var bearer = "Bearer " + os.Getenv("SOLANABEACH_API_KEY")

	// Create a new request using http
	req, err := http.NewRequest("GET", q, nil)
	req.Header.Add("Authorization", bearer)
	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Fail to get tx from rpc:%s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == 429 || resp.StatusCode == 524 {
		time.Sleep(time.Second)
		return client.get(query)
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Fail to get tx from rpc:Status :%s", resp.Status)
	}

	var bz []byte
	bz, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return bz, nil
}
