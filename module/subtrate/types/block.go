package types

type Block struct {
	Code int `json:"code"`
	Data struct {
		AccountDisplay struct {
			AccountIndex  string      `json:"account_index"`
			Address       string      `json:"address"`
			Display       string      `json:"display"`
			Identity      bool        `json:"identity"`
			Judgements    interface{} `json:"judgements"`
			Parent        string      `json:"parent"`
			ParentDisplay string      `json:"parent_display"`
		} `json:"account_display"`
		BlockNum       int           `json:"block_num"`
		BlockTimestamp int           `json:"block_timestamp"`
		EventCount     int           `json:"event_count"`
		Events         []interface{} `json:"events"`
		Extrinsics     []struct {
			AccountDisplay     interface{} `json:"account_display"`
			AccountID          string      `json:"account_id"`
			AccountIndex       string      `json:"account_index"`
			BlockNum           int         `json:"block_num"`
			BlockTimestamp     int         `json:"block_timestamp"`
			CallModule         string      `json:"call_module"`
			CallModuleFunction string      `json:"call_module_function"`
			ExtrinsicHash      string      `json:"extrinsic_hash"`
			ExtrinsicIndex     string      `json:"extrinsic_index"`
			Fee                string      `json:"fee"`
			Nonce              int         `json:"nonce"`
			Params             string      `json:"params"`
			Signature          string      `json:"signature"`
			Success            bool        `json:"success"`
		} `json:"extrinsics"`
		ExtrinsicsCount int    `json:"extrinsics_count"`
		ExtrinsicsRoot  string `json:"extrinsics_root"`
		Finalized       bool   `json:"finalized"`
		Hash            string `json:"hash"`
		Logs            []struct {
			BlockNum   int    `json:"block_num"`
			Data       string `json:"data"`
			LogIndex   string `json:"log_index"`
			LogType    string `json:"log_type"`
			OriginType string `json:"origin_type"`
		} `json:"logs"`
		ParentHash        string `json:"parent_hash"`
		SpecVersion       int    `json:"spec_version"`
		StateRoot         string `json:"state_root"`
		Validator         string `json:"validator"`
		ValidatorIndexIds string `json:"validator_index_ids"`
		ValidatorName     string `json:"validator_name"`
	} `json:"data"`
	Message     string `json:"message"`
	GeneratedAt int    `json:"generated_at"`
}

func (v *Block) SubscanApi() {}
