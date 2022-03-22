package input

type BigChainType struct{
	ChainType string `json:"chain_type"`
	Details []IndividualChain `json:"details"`
}

type IndividualChain struct{
	ChainName string `json:"chain_name"`
	Validators []ValidatorDetail `json:"validators"`
	FundHoldingAccount []string `json:"fund_holding_account"`
}

type ValidatorDetail struct{
	ValidatorAddress string `json:"validator_address"`
	SelfDelegationAddress string `json:"self_delegation_address"`
}

type Data struct{
	Data []BigChainType `json:"data"`
}