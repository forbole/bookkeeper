package utils

import (
	client "github.com/forbole/bookkeeper/module/solana/client"

)

// GetDelegatorAddresses the addresses associated to the pubkey that 
// delegated to the validator specified
func GetSelfDelegatorAddresses(pubKey string,validatorIdentity string,client *client.SolanaBeachClient)([]string,error){
	accounts,err:=client.GetStakeAccounts(pubKey)
	if err!=nil{
		return nil,err
	}

	var addresses []string
	for _,account:=range accounts{
		if account.Data.Stake.Delegation.ValidatorInfo.IdentityPubkey==validatorIdentity{
			addresses=append(addresses,account.Pubkey.Address)
		}
	}
	return addresses,nil

}

func GetDelegatorAddress(pubKey string,client *client.SolanaBeachClient)([]string,error){
	accounts,err:=client.GetStakeAccounts(pubKey)
	if err!=nil{
		return nil,err
	}

	var addresses []string
	for _,account:=range accounts{
			addresses=append(addresses,account.Pubkey.Address)
	}
	return addresses,nil
}