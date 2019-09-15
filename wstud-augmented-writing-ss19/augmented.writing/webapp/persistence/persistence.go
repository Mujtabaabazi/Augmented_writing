epackage persistence

import (
	"database/sql"
	"fmt"
)

var db *sql.DB // other packages shoulD access this via "business.Connection()"

func Init() {
	connectToDatabase()
}

func Connection() *sql.DB {

	if db == nil {
		connectToDatabase()
	}
	return db
}

/**
Implement below
**/

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "mujtaba"
	dbname   = "augmented_writing"
)

func connectToDatabase() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if err != nil {
		panic(err)
	}
}
