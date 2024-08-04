package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"os"
)

type Db struct {
	conn *pgx.Conn
	ctx  context.Context
}

func (db *Db) InitDbConnection(ctx context.Context) {

	dbUrl := os.Getenv("DATABASE_URL")
	fmt.Printf("Connecting to database at %s\n", dbUrl)

	conn, err := pgx.Connect(ctx, dbUrl)
	if err != nil {
		panic(err.Error())
	}
	db.conn = conn
	db.ctx = ctx
}

func (db *Db) Close() {
	err := db.conn.Close(db.ctx)
	if err != nil {
		panic(err.Error())
	}
}
