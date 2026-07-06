package database

import (
	"os"
    "database/sql"

    _ "github.com/mattn/go-sqlite3"
)
var DB *sql.DB
func InitDb() (error) {
	var err error
	DB, err = sql.Open("sqlite3", "DB.db")
	if err != nil {
		return err
	}

	schema, err := os.ReadFile("database/schema.sql")
	if err != nil {
		return err
	}
	_, err = DB.Exec(string(schema))
	if err != nil {
		DB.Close()
		return err
	}
	return err
}