package business

import (
	"database/sql"

	"github.com/Augmented_writing/wstud-augmented-writing-ss19/augmented.writing/webapp/persistence"
)

// define all SQL queries and create 1 func for each call

const (
	SelectUser = `SELECT ... FROM ... WHERE ...`
)

var db *sql.DB

// DummyFunc connects to the datanase by one line..
func DummyFunc() *sql.DB {

	db = persistence.Connection()

	return db
}
