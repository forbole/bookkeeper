package types

import (
	"fmt"
	"time"
)

// ValidatorStatus table is the validator status at that time
type ValidatorStatus struct{
	Time time.Time
	ChainId string
	ValidatorDelegationCount float64
	ValidatorCommissionRate float64
	TotalVotingPower float64
	ValidatorVotingPowerRanking float64
	StakeAmount float64
	ValidatorVotingPower float64
}

func NewValidatorStatus(time time.Time,chainId string, validatorDelegatorCount float64,
						validatorCommissionRate float64,
						totalVotingPower float64, validatorVotingPowerRanking float64,
						stakeAmount float64,validatorVotingPower float64)ValidatorStatus{
							return ValidatorStatus{
								Time: time,
								ChainId:chainId,
								ValidatorDelegationCount: validatorDelegatorCount,
								ValidatorCommissionRate:validatorCommissionRate,
								TotalVotingPower:totalVotingPower,
								ValidatorVotingPowerRanking :validatorVotingPowerRanking,
								StakeAmount:stakeAmount,
								ValidatorVotingPower: validatorVotingPower,
							}
						}

type ValidatorStatusTable []ValidatorStatus

func (v ValidatorStatusTable) GetCSV()string{
	outputcsv:="time,chain_id,ValidatorDelegationCount,ValidatorCommissionRate,TotalVotingPower,ValidatorVotingPowerRanking,StakeAmount,ValidatorVotingPower\n"
	for _, b := range v {
		outputcsv += fmt.Sprintf("%s,%s,%f,%f,%f,%f,%f,%f\n",
			b.Time,b.ChainId,b.ValidatorDelegationCount,b.ValidatorCommissionRate,b.TotalVotingPower,b.ValidatorVotingPowerRanking,b.StakeAmount,b.TotalVotingPower,
	)
	}
	return outputcsv
}
