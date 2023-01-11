package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// TODO: Move to .env file
var (
	host     = "database"
	port     = "5432"
	user     = "goinglive"
	password = "goinglive"
	dbname   = "goinglive"
)

var db *sql.DB

func GetDB() *sql.DB {
	if db != nil {
		return db
	}
	db = initDBPool()
	return db
}

func initDBPool() *sql.DB {

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		user,
		password,
		dbname,
	)

	db, err := sql.Open("postgres", dsn)

	if err != nil {
		panic(err)
	}

	return db
}
