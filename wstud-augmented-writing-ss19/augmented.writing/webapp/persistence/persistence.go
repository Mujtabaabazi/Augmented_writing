package persistence

import (
	"fmt"
	"github.com/jinzhu/gorm"
	// importing to register postgres driver
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB // other packages shoulD access this via "business.Connection()"

func Init() {
	connectToDatabase()
}

func Connection() *gorm.DB {

	if DB == nil {
		connectToDatabase()
	}
	return DB
}

/**
Implement below
**/

const (
	host     = "host.docker.internal"
	port     = 30532
	user     = "postgres"
	password = "postgres"
	dbname   = "test"
	dialect  = "postgres"
	//host     = "localhost"
	//port     = 5432
	//user     = "postgres"
	//password = "mujtaba"
	//dbname   = "augmented_writing"
)

func connectToDatabase() {
	//psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	//	"password=%s dbname=%s sslmode=disable",
	//	host, port, user, password, dbname)
	connStr := fmt.Sprintf(
		"postgres://%v:%v@%v:%v/%v?sslmode=disable",
		user,
		password,
		host,
		port,
		dbname,
	)
	db, err := gorm.Open(dialect, connStr)
	if err != nil {
		panic(err)
	}
	DB = db
}
