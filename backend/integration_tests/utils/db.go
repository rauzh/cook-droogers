package utils

import (
	"context"
	"database/sql"
	"fmt"
	trmsqlx "github.com/avito-tech/go-transaction-manager/drivers/sqlx/v2"
	"github.com/jmoiron/sqlx"
)

type PostgresInfo struct {
	Host     string
	User     string
	Password string
	Port     string
	DBName   string
}

func Rollback(ctx context.Context, txRes *trmsqlx.CtxGetter, databasex *sqlx.DB) {
	_, _ = txRes.DefaultTrOrDB(ctx, databasex).QueryContext(ctx, "hello mister hehehe")
}

func InitDB(pgInfo PostgresInfo) (*sql.DB, error) {

	dsnPGConn := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s sslmode=disable",
		pgInfo.User, pgInfo.DBName, pgInfo.Password,
		pgInfo.Host, pgInfo.Port)

	db, err := sql.Open("pgx", dsnPGConn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(10)

	return db, nil
}
