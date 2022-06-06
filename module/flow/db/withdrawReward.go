package db

import (
    "github.com/jmoiron/sqlx"

)
const (
	host     = "45.79.148.99"
	port     = 5432
	user     = "postgres"
	password = ""
	dbname   = "flowjuno"
  )

type FlowDb struct{
	sql *sqlx.DB
}

func build(connStr string)(*FlowDb,error){
	postgresDb, err := sqlx.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return &FlowDb{sql:postgresDb},nil
}

// GetWithdrawReward get the withdraw reward from the flowjuno db directly...
func (db *FlowDb)GetWithdrawReward(payer string)([]heightValue,error){
	stmt:=`select height,value from transaction left join event on transaction.transaction_id = event.transaction_id 
	where payer='$1' and type='A.8624b52f9ddcd04a.FlowIDTableStaking.RewardTokensWithdrawn'`

	var heightValue []heightValue
	err:=db.sql.Select(&heightValue,stmt,payer)
	if err!=nil{
		return nil,err
	}

	return heightValue,nil
}



type heightValue struct{
	height int64
	value string
}