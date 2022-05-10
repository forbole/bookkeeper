package types

type ValidatorStat struct {
	Status string `json:"status"`
	Data   struct {
		ResultType string `json:"resultType"`
		Result     []struct {
			Metric struct {
				Name        string `json:"__name__"`
				ChainID     string `json:"chain_id"`
				Denom       string `json:"denom"`
				Environment string `json:"environment"`
				Host        string `json:"host"`
				Instance    string `json:"instance"`
				Job         string `json:"job"`
			} `json:"metric"`
			Value []interface{} `json:"value"`
		} `json:"result"`
	} `json:"data"`
}

type ValidatorDelegationCount struct {
	Status string `json:"status"`
	Data   struct {
		ResultType string `json:"resultType"`
		Result     []struct {
			Metric struct {
				Name             string `json:"__name__"`
				ChainID          string `json:"chain_id"`
				Environment      string `json:"environment"`
				Host             string `json:"host"`
				Instance         string `json:"instance"`
				Job              string `json:"job"`
				ValidatorAddress string `json:"validator_address"`
			} `json:"metric"`
			Value []interface{} `json:"value"`
		} `json:"result"`
	} `json:"data"`
}

type StakeAmount struct {
	Status string `json:"status"`
	Data   struct {
		ResultType string `json:"resultType"`
		Result     []struct {
			Metric struct {
				Name             string `json:"__name__"`
				ChainID          string `json:"chain_id"`
				DelegatorAddress string `json:"delegator_address"`
				Denom            string `json:"denom"`
				Environment      string `json:"environment"`
				Host             string `json:"host"`
				Instance         string `json:"instance"`
				Job              string `json:"job"`
				ValidatorAddress string `json:"validator_address"`
			} `json:"metric"`
			Value []interface{} `json:"value"`
		} `json:"result"`
	} `json:"data"`
}

type ValidatorCommissionRate struct {
	Status string `json:"status"`
	Data   struct {
		ResultType string `json:"resultType"`
		Result     []struct {
			Metric struct {
				Name             string `json:"__name__"`
				ChainID          string `json:"chain_id"`
				Environment      string `json:"environment"`
				Host             string `json:"host"`
				Instance         string `json:"instance"`
				Job              string `json:"job"`
				ValidatorAddress string `json:"validator_address"`
			} `json:"metric"`
			Value []interface{} `json:"value"`
		} `json:"result"`
	} `json:"data"`
}

type ValidatorVotingPower struct {
	Status string `json:"status"`
	Data   struct {
		ResultType string `json:"resultType"`
		Result     []struct {
			Metric struct {
				Name             string `json:"__name__"`
				ChainID          string `json:"chain_id"`
				Denom            string `json:"denom"`
				Environment      string `json:"environment"`
				Host             string `json:"host"`
				Instance         string `json:"instance"`
				Job              string `json:"job"`
				ValidatorAddress string `json:"validator_address"`
			} `json:"metric"`
			Value []interface{} `json:"value"`
		} `json:"result"`
	} `json:"data"`
}

type TotalVotingPower struct {
	Status string `json:"status"`
	Data   struct {
		ResultType string `json:"resultType"`
		Result     []struct {
			Metric struct {
				Name        string `json:"__name__"`
				ChainID     string `json:"chain_id"`
				Denom       string `json:"denom"`
				Environment string `json:"environment"`
				Host        string `json:"host"`
				Instance    string `json:"instance"`
				Job         string `json:"job"`
			} `json:"metric"`
			Value []interface{} `json:"value"`
		} `json:"result"`
	} `json:"data"`
}

type ValidatorVotingPowerRanking struct {
	Status string `json:"status"`
	Data   struct {
		ResultType string `json:"resultType"`
		Result     []struct {
			Metric struct {
				Name        string `json:"__name__"`
				ChainID     string `json:"chain_id"`
				Environment string `json:"environment"`
				Host        string `json:"host"`
				Instance    string `json:"instance"`
				Job         string `json:"job"`
			} `json:"metric"`
			Value []interface{} `json:"value"`
		} `json:"result"`
	} `json:"data"`
}
