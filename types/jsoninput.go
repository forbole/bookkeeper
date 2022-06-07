package types

type Subtrate struct {
	ChainName string   `json:"chain_name"`
	Address   []string `json:"address"`
	Denom     []Denom  `json:"denom"`
}

type Denom struct {
	Denom    string `json:"denom"`
	Exponent int    `json:"exponent"`
	CoinId   string `json:"coin_id"`
	Cointype string `json:"cointype"`
}

type CosmosDetails struct {
	ChainName          string            `json:"chain_name"`
	Denom              []Denom           `json:"denom"`
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
	Chains   []CosmosDetails `json:"chains"`
	Subtrate []Subtrate      `json:"subtrate"`

	EmailDetails EmailDetails `json:"email_details"`
	Prometheus   string       `json:"prometheus"`
	Flow         Flow         `json:"flow"`
	VsCurrency   string       `json:"vs_currency"`
	Period       Period       `json:"period"`
}

// Period get the unix time period from and until date
type Period struct {
	From int64 `json:"from"`
	To   int64 `json:"to"`
}

//host=%s port=%d dbname=%s user=%s sslmode=%s search_path=%s
type Database struct{
	Host string `json:"host"`
	Port int `json:"port"`
	DbName string `json:"db_name"`
	User string `json:"user"`
	SSLMode string `json:"ssl_mode"`
	SearchPath string `json:"search_path"`
	Password string `json:"password"`
}

type Flow struct {
	FlowJuno     string   `json:"flowjuno"`
	FlowEndpoint string   `json:"flow_endpoint"`
	Addresses    []string `json:"addresses"`
	Denom        string   `json:"denom"`
	Exponent     int      `json:"exponent"`
	LastSpork    int      `json:"last_spork"`
	Db Database `json:"database"`
}
