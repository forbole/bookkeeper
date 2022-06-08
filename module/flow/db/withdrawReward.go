package db

import (
	"context"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"

	"github.com/forbole/bookkeeper/types"
	"github.com/jmoiron/sqlx"
)

type FlowDb struct{
	sql *sqlx.DB
}

func Build(dbSpec types.Database)(*FlowDb,error){
	log.Trace().Str("module", "flow").Msg("build database")

	fmt.Println(dbSpec.DbName)
	fmt.Println(dbSpec.User)
	fmt.Println(dbSpec.Host)


	connstr:=fmt.Sprintf("host=%s port=%d dbname=%s user=%s sslmode=%s search_path=%s password=%s",
	dbSpec.Host,dbSpec.Port,dbSpec.DbName,dbSpec.User,dbSpec.SSLMode,dbSpec.SearchPath,dbSpec.Password)

	postgresDb, err := sqlx.Open("postgres", connstr)
	if err != nil {
		return nil, err
	}

	if err = postgresDb.Ping(); err != nil {
        return nil, err
    }

	return &FlowDb{sql:postgresDb},nil
}

// GetWithdrawReward get the withdraw reward from the flowjuno db directly...
func (db *FlowDb)GetWithdrawReward(payer string)([]HeightValue,error){
	log.Trace().Str("module", "flow").Msg("get reward from db")

	
	stmt:=`select transaction.transaction_id as transaction_id,event.height as height,value from transaction left join event on transaction.transaction_id = event.transaction_id 
	where payer=$1 and type='A.8624b52f9ddcd04a.FlowIDTableStaking.RewardTokensWithdrawn' order by event.height desc`

	
	var heightValue []HeightValue
	err:=db.sql.SelectContext(context.Background(),&heightValue,stmt,payer)

	//err:=db.sql.Select(&heightValue,stmt,payer)
	if err!=nil{
		return nil,err
	}

	return heightValue,nil
}



type HeightValue struct{
	Height int64 `db:"height"`
	Value string `db:"value"`
	TransactionId string `db:"transaction_id"`
}