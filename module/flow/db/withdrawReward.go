package db

import (
	"fmt"

	"github.com/forbole/bookkeeper/types"
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

func Build(dbSpec types.Database)(*FlowDb,error){
	connstr:=fmt.Sprintf("host=%s port=%d dbname=%s user=%s sslmode=%s search_path=%s password=%s",
	dbSpec.Host,dbSpec.Port,dbSpec.DbName,dbSpec.User,dbSpec.SSLMode,dbSpec.SearchPath,dbSpec.Password)

	postgresDb, err := sqlx.Open("postgres", connstr)
	if err != nil {
		return nil, err
	}
	return &FlowDb{sql:postgresDb},nil
}

// GetWithdrawReward get the withdraw reward from the flowjuno db directly...
func (db *FlowDb)GetWithdrawReward(payer string)([]HeightValue,error){
	stmt:=`select transaction_id,event.height as height,value from transaction left join event on transaction.transaction_id = event.transaction_id 
	where payer='$1' and type='A.8624b52f9ddcd04a.FlowIDTableStaking.RewardTokensWithdrawn'`

	var heightValue []HeightValue
	err:=db.sql.Select(&heightValue,stmt,payer)
	if err!=nil{
		return nil,err
	}

	return heightValue,nil
}



type HeightValue struct{
	Height int64
	Value string
	TransactionId string
}