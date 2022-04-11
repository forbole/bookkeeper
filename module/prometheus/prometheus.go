package prometheus

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	promtypes "github.com/forbole/bookkeeper/module/prometheus/types"
	"github.com/forbole/bookkeeper/types"
)

func GetValidatorDetailsFromPrometheus(endpoint string)(error){
	validatorDelegationCount,err := getValidatorDelegationCount(endpoint)
	if err!=nil{
		return err
	}
	stakeAmount,err := getStakeAmount(endpoint)
	if err!=nil{
		return err
	}
	totalVotingPower,err := getTotalVotingPower(endpoint)
	if err!=nil{
		return err
	}
	validatorVotingPowerRanking,err := getValidatorVotingPowerRanking(endpoint)
	if err!=nil{
		return err
	}
	validatorVotingPower,err := getValidatorVotingPower(endpoint)
	if err!=nil{
		return err
	}
	validatorCommissionRate,err := getValidatorCommissionRate(endpoint)
	if err!=nil{
		return err
	}

	for _,r:=range validatorDelegationCount.Data.Result{
		chain:=r.Metric.ChainID
		validatorAddress:=r.Metric.ValidatorAddress
		commissionRate:=0
		timestampUnix:=fmt.Sprint(r.Value[0])
		time:=strings.Split(timestampUnix,".")
		timestamp,err:=strconv.ParseFloat(timestampUnix,64)
		if err!=nil{
			return err
		}
		timestampString,err:=strconv.ParseFloat(timestampUnix,64)
		if err!=nil{
			return err
		}
		time.Unix

		// search for the same chain-id and same validator
		for _,commissionRate:=range validatorCommissionRate.Data.Result{
			if commissionRate.Metric.ChainID==chain && 
			validatorAddress==commissionRate.Metric.ValidatorAddress{
				rate:=fmt.Sprint(commissionRate.Value[1])
				commissionRate,err:=strconv.ParseFloat(rate,64)
				if err!=nil{
					return err
				}
			}
		}



	}

return nil
}

func getValidatorDelegationCount(endpoint string)(*promtypes.ValidatorDelegationCount,error){
	query := fmt.Sprintf(`%s/prometheus/api/v1/query?query=%s`,
	endpoint,"validator_delegation_count" )
	fmt.Println(query)
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

	return &txSearchRes,nil

}

func getStakeAmount(endpoint string)(*promtypes.StakeAmount,error){
	query := fmt.Sprintf(`%s/prometheus/api/v1/query?query=%s`,
	endpoint,"stake_amount" )
	fmt.Println(query)
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

	return &txSearchRes,nil

}

func getTotalVotingPower(endpoint string)(*promtypes.TotalVotingPower,error){
	query := fmt.Sprintf(`%s/prometheus/api/v1/query?query=%s`,
	endpoint,"total_voting_power" )
	fmt.Println(query)
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

	return &txSearchRes,nil

}

func getValidatorVotingPowerRanking(endpoint string)(*promtypes.ValidatorVotingPowerRanking,error){
	query := fmt.Sprintf(`%s/prometheus/api/v1/query?query=%s`,
	endpoint,"validator_voting_power_ranking" )
	fmt.Println(query)
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

	return &txSearchRes,nil

}

//validator_voting_power
func getValidatorVotingPower(endpoint string)(*promtypes.ValidatorVotingPower,error){
	query := fmt.Sprintf(`%s/prometheus/api/v1/query?query=%s`,
	endpoint,"validator_voting_power" )
	fmt.Println(query)
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

	return &txSearchRes,nil
}

//validator_commission_rate
func getValidatorCommissionRate(endpoint string)(*promtypes.ValidatorCommissionRate,error){
	query := fmt.Sprintf(`%s/prometheus/api/v1/query?query=%s`,
	endpoint,"validator_commission_rate" )
	fmt.Println(query)
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

	return &txSearchRes,nil
}