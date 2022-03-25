package distribution

func BankSend(api string, address string, limit int, page int)error{
	resp,err:=http.Get(fmt.Sprintf("%s/txs?message.module=distribution&transfer.sender=%s&limit=%d&page=%d",
		api,address,limit,page))
	if err!=nil{
		return err
	}

	defer resp.Body.Close()

	bz, err := ioutil.ReadAll(resp.Body)

	var bankreceipt types.BankRecipientRPC
	err=json.Unmarshal(bz,&bankreceipt)
	if err!=nil{
		return err
	}
	
	return nil
}