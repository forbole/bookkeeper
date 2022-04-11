package types

import (
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
}

func NewValidatorStatus(time time.Time,chainId string, validatorDelegatorCount float64,
						validatorCommissionRate float64,
						totalVotingPower float64, validatorVotingPowerRanking float64,
						validatorVotingPowerRaking float64, stakeAmount float64)ValidatorStatus{
							return ValidatorStatus{
								Time: time,
								ValidatorDelegationCount: validatorDelegatorCount,
								ValidatorCommissionRate:validatorCommissionRate,
								TotalVotingPower:totalVotingPower,
								ValidatorVotingPowerRanking :validatorVotingPowerRanking,
								StakeAmount:stakeAmount,

							}
						}