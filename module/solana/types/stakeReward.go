package types

type StakeReward struct {
	Epoch         int     `json:"epoch"`
	EffectiveSlot int     `json:"effectiveSlot"`
	Amount        int     `json:"amount"`
	PostBalance   int64   `json:"postBalance"`
	PercentChange float64 `json:"percentChange"`
	Apr           float64 `json:"apr"`
	Timestamp     int     `json:"timestamp"`
}