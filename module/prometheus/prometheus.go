package prometheus

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	v040 "github.com/cosmos/cosmos-sdk/x/genutil/legacy/v040"
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

	var validatorStatus []types.ValidatorStatus
	for _,r:=range validatorDelegationCount.Data.Result{
		chain:=r.Metric.ChainID
		validatorAddress:=r.Metric.ValidatorAddress
		delegationCount:=r.Value[0].(float64)

		commissionRate:=float64(0)
		selfStake:=float64(0)
		totalvp:=float64(0)
		vpRanking:=float64(0)
		vp:=float64(0)


		timestampUnix:=fmt.Sprint(r.Value[0])
		timestamp:=strings.Split(timestampUnix,".")
		timestampint,err:=strconv.ParseInt(timestamp[0],10,64)
		if err!=nil{
			return err
		}
		timeStampReal:=time.Unix(timestampint,0)

		// search for the same chain-id and same validator
		for _,result:=range validatorCommissionRate.Data.Result{
			if result.Metric.ChainID==chain {
				commissionRate,ok:=result.Value[1].(float64)
				if !ok{
					return fmt.Errorf("CommissionRate is not float64")
				}
			}
		}

		for _,result:=range stakeAmount.Data.Result{
			if result.Metric.ChainID==chain {
				selfStake,ok:=result.Value[1].(float64)
				if !ok{
					return fmt.Errorf("selfStake is not float64")
				}
			}
		}

		for _,result:=range totalVotingPower.Data.Result{
			if result.Metric.ChainID==chain {
				totalvp,ok:=result.Value[1].(float64)
				if !ok{
					return fmt.Errorf("selfStake is not float64")
				}
			}
		}
		for _,result:=range validatorVotingPowerRanking.Data.Result{
			if result.Metric.ChainID==chain {
				vpRanking,ok:=result.Value[1].(float64)
				if !ok{
					return fmt.Errorf("selfStake is not float64")
				}
				vpRanking=vpRanking
			}
		}

		for _,result:=range validatorVotingPower.Data.Result{
			if result.Metric.ChainID==chain {
				vpreslut,ok:=result.Value[1].(float64)
				if !ok{
					return fmt.Errorf("selfStake is not float64")
				}
				vp=vpreslut
			}
		}

		validatorStatus=append(validatorStatus,types.NewValidatorStatus(timeStampReal,
			chain,delegationCount,commissionRate,totalvp,vp,float64(vpRanking),float64(selfStake)) )

		
	}

return nil
}

func GetValue(v promtypes.ValidatorStat,chain string)(float64,error){
	value:=float64(0)
	for _,result:=range v.Data.Result{
		if result.Metric.ChainID==chain {
			rate:=fmt.Sprint(result.Value[1])
			val,err:=strconv.ParseFloat(rate,64)
			if err!=nil{
				return 0,err
			}
			value=val
		}
	}
	return value,nil
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