package business

import (
	"database/sql"

	"wstud-augmented-writing-ss19/augmented.writing/webapp/persistence"
)

// define all SQL queries and create 1 func for each call

const (
	SelectUser = `SELECT ... FROM ... WHERE ...`
)

var db *sql.DB

// DummyFunc connects to the datanase by one line..
func DummyFunc() {

	db = persistence.Connection()

	return
}
