package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	HOST          = "db"
	DATABASE      = "rinhadb"
	USER          = "postgres"
	PASSWORD      = "postgres"
	MAX_OPEN_CONN = 20
	MAX_IDLE_CONN = 20
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func ConnectToDatabase() *sql.DB {
	var connectionString string = fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", HOST, USER, PASSWORD, DATABASE)
	db, err := sql.Open("postgres", connectionString)
	checkError(err)

	err = db.Ping()
	checkError(err)
	db.SetMaxOpenConns(MAX_OPEN_CONN)
	db.SetMaxOpenConns(MAX_IDLE_CONN)
	CreateTable(db)
	return db
}

func CreateTable(db *sql.DB) {
	create_table_query := `CREATE TABLE IF NOT EXISTS people (
			id uuid NOT NULL,
			nickname varchar(32) PRIMARY KEY NOT NULL,
			"name" varchar(100) NOT NULL,
			birthday date NULL,
			stack text NULL
	);`
	_, err := db.Exec(create_table_query)
	if err != nil {
		panic(err.Error())
	}
}
