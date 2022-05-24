package types
type RewardSlash struct {
	Code        int    `json:"code"`
	Message     string `json:"message"`
	GeneratedAt int    `json:"generated_at"`
	Data        struct {
		Count int `json:"count"`
		List  []struct {
			EventIndex    string `json:"event_index"`
			BlockNum      int    `json:"block_num"`
			ExtrinsicIdx  int    `json:"extrinsic_idx"`
			ModuleID      string `json:"module_id"`
			EventID       string `json:"event_id"`
			Params        string `json:"params"`
			ExtrinsicHash string `json:"extrinsic_hash"`
			EventIdx      int    `json:"event_idx"`
		} `json:"list"`
	} `json:"data"`
}

func (v *RewardSlash) SubscanApi(){

}

