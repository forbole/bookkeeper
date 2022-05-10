package prometheus

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	promtypes "github.com/forbole/bookkeeper/module/prometheus/types"
	"github.com/forbole/bookkeeper/types/tabletypes"
)

func GetValidatorDetailsFromPrometheus(endpoint string) (tabletypes.ValidatorStatusTable, error) {
	validatorDelegationCount, err := getValidatorDelegationCount(endpoint)
	if err != nil {
		return nil, err
	}
	stakeAmount, err := getStakeAmount(endpoint)
	if err != nil {
		return nil, err
	}
	totalVotingPower, err := getTotalVotingPower(endpoint)
	if err != nil {
		return nil, err
	}
	validatorVotingPowerRanking, err := getValidatorVotingPowerRanking(endpoint)
	if err != nil {
		return nil, err
	}
	validatorVotingPower, err := getValidatorVotingPower(endpoint)
	if err != nil {
		return nil, err
	}
	validatorCommissionRate, err := getValidatorCommissionRate(endpoint)
	if err != nil {
		return nil, err
	}

	var validatorStatus []tabletypes.ValidatorStatus
	for _, r := range validatorDelegationCount.Data.Result {
		chain := r.Metric.ChainID
		////fmt.Println(chain)
		val, ok := r.Value[1].(string)
		if !ok {
			return nil, fmt.Errorf("validatorCommissionRate is not string")
		}
		value, err := strconv.ParseFloat(val, 64)
		if err != nil {
			return nil, err
		}
		delegationCount := value

		commissionRate := float64(0)
		selfStake := float64(0)
		totalvp := float64(0)
		vpRanking := float64(0)
		vp := float64(0)

		timestampUnix, ok := r.Value[0].(float64)
		if !ok {
			return nil, fmt.Errorf("Timestamp is not float64")
		}
		if err != nil {
			return nil, err
		}

		//fmt.Println(int64(timestampUnix))
		timeStampReal := time.Unix(int64(timestampUnix), 0)

		// search for the same chain-id and same validator
		for _, result := range validatorCommissionRate.Data.Result {
			if result.Metric.ChainID == chain {
				val, ok := result.Value[1].(string)
				if !ok {
					return nil, fmt.Errorf("validatorCommissionRate is not string")
				}
				value, err := strconv.ParseFloat(val, 64)
				if err != nil {
					return nil, err
				}

				commissionRate = value
			}
		}

		for _, result := range stakeAmount.Data.Result {
			if result.Metric.ChainID == chain {
				val, ok := result.Value[1].(string)
				if !ok {
					return nil, fmt.Errorf("validatorCommissionRate is not string")
				}
				value, err := strconv.ParseFloat(val, 64)
				if err != nil {
					return nil, err
				}
				selfStake = value
			}
		}

		for _, result := range totalVotingPower.Data.Result {
			if result.Metric.ChainID == chain {
				val, ok := result.Value[1].(string)
				if !ok {
					return nil, fmt.Errorf("validatorCommissionRate is not string")
				}
				value, err := strconv.ParseFloat(val, 64)
				if err != nil {
					return nil, err
				}
				totalvp = value
			}
		}
		for _, result := range validatorVotingPowerRanking.Data.Result {
			if result.Metric.ChainID == chain {
				val, ok := result.Value[1].(string)
				if !ok {
					return nil, fmt.Errorf("validatorCommissionRate is not float64")
				}
				value, err := strconv.ParseFloat(val, 64)
				if err != nil {
					return nil, err
				}
				vpRanking = value
			}
		}

		for _, result := range validatorVotingPower.Data.Result {
			if result.Metric.ChainID == chain {
				val, ok := result.Value[1].(string)
				if !ok {
					return nil, fmt.Errorf("validatorCommissionRate is not string")
				}
				value, err := strconv.ParseFloat(val, 64)
				if err != nil {
					return nil, err
				}
				vp = value
			}
		}

		status := tabletypes.NewValidatorStatus(timeStampReal,
			r.Metric.ChainID, delegationCount, commissionRate, totalvp, vpRanking, selfStake, vp)
		//fmt.Println(chain)

		validatorStatus = append(validatorStatus, status)

	}

	return validatorStatus, nil
}

func getValidatorDelegationCount(endpoint string) (*promtypes.ValidatorDelegationCount, error) {
	query := fmt.Sprintf(`%s/prometheus/api/v1/query?query=%s`,
		endpoint, "validator_delegation_count")
	//fmt.Println(query)
	resp, err := http.Get(query)
	if err != nil {
		return nil, fmt.Errorf("Fail to get tx from rpc:%s", err)
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Fail to get tx from rpc:Status code:%d", resp.StatusCode)
	}

	defer resp.Body.Close()

	bz, err := io.ReadAll(resp.Body)

	var txSearchRes promtypes.ValidatorDelegationCount
	err = json.Unmarshal(bz, &txSearchRes)
	if err != nil {
		return nil, fmt.Errorf("Fail to marshal:%s", err)
	}

	return &txSearchRes, nil

}

func getStakeAmount(endpoint string) (*promtypes.StakeAmount, error) {
	query := fmt.Sprintf(`%s/prometheus/api/v1/query?query=%s`,
		endpoint, "stake_amount")
	//fmt.Println(query)
	resp, err := http.Get(query)
	if err != nil {
		return nil, fmt.Errorf("Fail to get tx from rpc:%s", err)
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Fail to get tx from rpc:Status code:%d", resp.StatusCode)
	}

	defer resp.Body.Close()

	bz, err := io.ReadAll(resp.Body)

	var txSearchRes promtypes.StakeAmount
	err = json.Unmarshal(bz, &txSearchRes)
	if err != nil {
		return nil, fmt.Errorf("Fail to marshal:%s", err)
	}

	return &txSearchRes, nil

}

func getTotalVotingPower(endpoint string) (*promtypes.TotalVotingPower, error) {
	query := fmt.Sprintf(`%s/prometheus/api/v1/query?query=%s`,
		endpoint, "total_voting_power")
	//fmt.Println(query)
	resp, err := http.Get(query)
	if err != nil {
		return nil, fmt.Errorf("Fail to get tx from rpc:%s", err)
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Fail to get tx from rpc:Status code:%d", resp.StatusCode)
	}

	defer resp.Body.Close()

	bz, err := io.ReadAll(resp.Body)

	var txSearchRes promtypes.TotalVotingPower
	err = json.Unmarshal(bz, &txSearchRes)
	if err != nil {
		return nil, fmt.Errorf("Fail to marshal:%s", err)
	}

	return &txSearchRes, nil

}

func getValidatorVotingPowerRanking(endpoint string) (*promtypes.ValidatorVotingPowerRanking, error) {
	query := fmt.Sprintf(`%s/prometheus/api/v1/query?query=%s`,
		endpoint, "validator_voting_power_ranking")
	//fmt.Println(query)
	resp, err := http.Get(query)
	if err != nil {
		return nil, fmt.Errorf("Fail to get tx from rpc:%s", err)
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Fail to get tx from rpc:Status code:%d", resp.StatusCode)
	}

	defer resp.Body.Close()

	bz, err := io.ReadAll(resp.Body)

	var txSearchRes promtypes.ValidatorVotingPowerRanking
	err = json.Unmarshal(bz, &txSearchRes)
	if err != nil {
		return nil, fmt.Errorf("Fail to marshal:%s", err)
	}

	return &txSearchRes, nil

}

//validator_voting_power
func getValidatorVotingPower(endpoint string) (*promtypes.ValidatorVotingPower, error) {
	query := fmt.Sprintf(`%s/prometheus/api/v1/query?query=%s`,
		endpoint, "validator_voting_power")
	//fmt.Println(query)
	resp, err := http.Get(query)
	if err != nil {
		return nil, fmt.Errorf("Fail to get tx from rpc:%s", err)
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Fail to get tx from rpc:Status code:%d", resp.StatusCode)
	}

	defer resp.Body.Close()

	bz, err := io.ReadAll(resp.Body)

	var txSearchRes promtypes.ValidatorVotingPower
	err = json.Unmarshal(bz, &txSearchRes)
	if err != nil {
		return nil, fmt.Errorf("Fail to marshal:%s", err)
	}

	return &txSearchRes, nil
}

//validator_commission_rate
func getValidatorCommissionRate(endpoint string) (*promtypes.ValidatorCommissionRate, error) {
	query := fmt.Sprintf(`%s/prometheus/api/v1/query?query=%s`,
		endpoint, "validator_commission_rate")
	//fmt.Println(query)
	resp, err := http.Get(query)
	if err != nil {
		return nil, fmt.Errorf("Fail to get tx from rpc:%s", err)
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Fail to get tx from rpc:Status code:%d", resp.StatusCode)
	}

	defer resp.Body.Close()

	bz, err := io.ReadAll(resp.Body)

	var txSearchRes promtypes.ValidatorCommissionRate
	err = json.Unmarshal(bz, &txSearchRes)
	if err != nil {
		return nil, fmt.Errorf("Fail to marshal:%s", err)
	}

	return &txSearchRes, nil
}
