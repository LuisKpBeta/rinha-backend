package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"

	_ "github.com/lib/pq"
)

const (
	HOST          = "db"
	DATABASE      = "rinhadb"
	USER          = "postgres"
	PASSWORD      = "postgres"
	MAX_OPEN_CONN = 30
	MAX_IDLE_CONN = 15
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func ConnectToDatabase() *pgxpool.Pool {
	var connectionString string = fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", HOST, USER, PASSWORD, DATABASE)

	poolConfig, err := pgxpool.ParseConfig(connectionString)
	checkError(err)

	db, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	checkError(err)

	err = db.Ping(context.Background())
	checkError(err)

	CreateTable(db)
	return db
}

func CreateTable(db *pgxpool.Pool) {
	create_table_query := `CREATE TABLE IF NOT EXISTS people (
			id uuid NOT NULL,
			nickname varchar(32) PRIMARY KEY NOT NULL,
			"name" varchar(100) NOT NULL,
			birthday date NULL,
			stacks text NULL
	);`
	_, err := db.Exec(context.Background(), create_table_query)
	if err != nil {
		panic(err.Error())
	}
}
