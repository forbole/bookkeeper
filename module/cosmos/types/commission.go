package types

type Commission struct {
	Commission struct {
		Commission []DenomAmount `json:"commission"`
	} `json:"commission"`
}

	
type Rewards struct {
	Rewards struct {
		Rewards []DenomAmount `json:"rewards"`
	} `json:"rewards"`
}

type DenomAmount struct {
		Denom  string `json:"denom"`
		Amount string `json:"amount"`
}