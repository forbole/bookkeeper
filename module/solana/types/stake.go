package types

type Stake struct {
	TotalPages int `json:"totalPages"`
	Data  []StakeData `json:"data"`
}

type StakeData struct{
		Pubkey struct {
			Address string `json:"address"`
		} `json:"pubkey"`
		Lamports int64 `json:"lamports"`
		Data     struct {
			State int `json:"state"`
			Meta  struct {
				RentExemptReserve int `json:"rent_exempt_reserve"`
				Authorized        struct {
					Staker struct {
						Address string `json:"address"`
					} `json:"staker"`
					Withdrawer struct {
						Address string `json:"address"`
					} `json:"withdrawer"`
				} `json:"authorized"`
				Lockup struct {
					UnixTimestamp int `json:"unix_timestamp"`
					Epoch         int `json:"epoch"`
					Custodian     struct {
						Address string `json:"address"`
					} `json:"custodian"`
				} `json:"lockup"`
			} `json:"meta"`
			Stake struct {
				Delegation struct {
					VoterPubkey struct {
						Address string `json:"address"`
					} `json:"voter_pubkey"`
					Stake              int64   `json:"stake"`
					ActivationEpoch    int     `json:"activation_epoch"`
					DeactivationEpoch  int64   `json:"deactivation_epoch"`
					WarmupCooldownRate float64 `json:"warmup_cooldown_rate"`
					ValidatorInfo      struct {
						Name            string `json:"name"`
						Website         string `json:"website"`
						IdentityPubkey  string `json:"identityPubkey"`
						KeybaseUsername string `json:"keybaseUsername"`
						Details         string `json:"details"`
						Image           string `json:"image"`
					} `json:"validatorInfo"`
				} `json:"delegation"`
				CreditsObserved int `json:"credits_observed"`
			} `json:"stake"`
		} `json:"data"`
}
