package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "database"
	port     = 5432
	user     = "goinglive"
	password = "goinglive"
	dbname   = "goinglive"
)

func GetDB() *sql.DB {

	postgresqlDbInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", postgresqlDbInfo)

	if err != nil {
		panic(err)
	}

	return db
}

var Client = GetDB()
