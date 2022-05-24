package types

type RewardSlash struct {
	Code        int    `json:"code"`
	Message     string `json:"message"`
	GeneratedAt int    `json:"generated_at"`
	Data        struct {
		Count int `json:"count"`
		List  []struct {
			Era            int    `json:"era"`
			Stash          string `json:"stash"`
			Account        string `json:"account"`
			ValidatorStash string `json:"validator_stash"`
			Amount         string `json:"amount"`
			BlockTimestamp int    `json:"block_timestamp"`
			EventIndex     string `json:"event_index"`
			ModuleID       string `json:"module_id"`
			EventID        string `json:"event_id"`
			SlashKton      string `json:"slash_kton"`
			ExtrinsicIndex string `json:"extrinsic_index"`
		} `json:"list"`
	} `json:"data"`
}

func (v *RewardSlash) SubscanApi() {}
