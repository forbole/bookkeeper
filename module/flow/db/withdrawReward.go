package db

import (
	"context"
	"fmt"
	"time"

	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"

	"github.com/forbole/bookkeeper/types"
	"github.com/jmoiron/sqlx"
)

type FlowDb struct{
	Sql *sqlx.DB
}

func Build(dbSpec types.Database)(*FlowDb,error){
	log.Trace().Str("module", "flow").Msg("build database")

	connstr:=fmt.Sprintf("host=%s port=%d dbname=%s user=%s sslmode=%s search_path=%s password=%s",
	dbSpec.Host,dbSpec.Port,dbSpec.DbName,dbSpec.User,dbSpec.SSLMode,dbSpec.SearchPath,dbSpec.Password)

	fmt.Println(connstr)
	postgresDb, err := sqlx.Open("postgres", connstr)
	if err != nil {
		return nil, fmt.Errorf("Cannot open connection:%s",err)
	}

	postgresDb.SetConnMaxLifetime(time.Minute * 5)
	postgresDb.SetConnMaxIdleTime(0)
	postgresDb.SetMaxIdleConns(0)

	if err = postgresDb.Ping(); err != nil {
        return nil, fmt.Errorf("Cannot ping:%s",err)
    }

	return &FlowDb{Sql:postgresDb},nil
}

// GetWithdrawReward get the withdraw reward from the flowjuno db directly...
func (db *FlowDb)GetWithdrawReward(payer string)([]HeightValue,error){
	log.Trace().Str("module", "flow").Msg("get reward from db")

	
	stmt:=`select transaction.transaction_id as transaction_id,event.height as height,value from transaction left join event on transaction.transaction_id = event.transaction_id 
	where payer=$1 and type='A.8624b52f9ddcd04a.FlowIDTableStaking.RewardTokensWithdrawn' order by event.height desc`

	
	var heightValue []HeightValue
	err:=db.Sql.SelectContext(context.Background(),&heightValue,stmt,payer)

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

func NewHeightValue(height int64,value,transactionId string)HeightValue{
	return HeightValue{
		Height: height,
		Value: value,
		TransactionId: transactionId,
	}
}