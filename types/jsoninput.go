package types

type BigChainType struct {
	ChainType string            `json:"chain_type"`
	Details   []IndividualChain `json:"details"`
}

type Denom struct{
	Denom string  `json:"denom"`
	Exponent           int `json:"exponent"`
	CoinId string `json:"coin_id"`
	Cointype string `json:"cointype"`
}

type IndividualChain struct {
	ChainName          string            `json:"chain_name"`
	Denom              []Denom            `json:"denom"`
	Validators         []ValidatorDetail `json:"validators"`
	FundHoldingAccount []string          `json:"fund_holding_account"`
	GrpcEndpoint       string            `json:"grpc_endpoint"`
	RpcEndpoint        string            `json:"rpc_endpoint"`
	LcdEndpoint        string            `json:"lcd_endpoint"`
}

type ValidatorDetail struct {
	ValidatorAddress      string `json:"validator_address"`
	SelfDelegationAddress string `json:"self_delegation_address"`
}

type EmailAccount struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Address  string `json:"address"`
	Host     string `json:"host"`
}

type EmailDetails struct {
	From    EmailAccount `json:"from"`
	To      []string     `json:"to"`
	Subject string       `json:"subject"`
	Details string       `json:"details"`
}

type Data struct {
	Chains       []BigChainType `json:"chains"`
	EmailDetails EmailDetails   `json:"email_details"`
	Prometheus   string         `json:"prometheus"`
	Flow Flow `json:"flow"`
	VsCurrency string `json:"vs_currency"`
	Period Period `json:"period"`
}

// Period get the unix time period from and until date
type Period struct{
	From int64 `json:"from"`
	To int64 `json:"to"`
}

type Flow struct{
	FlowJuno string `json:"flowjuno"`
	FlowEndpoint string `json:"flow_endpoint"`
	Addresses []string `json:"addresses"`
	Denom   string    `json:"denom"`
	Exponent int `json:"exponent"`
	LastSpork int `json:"last_spork"`
}