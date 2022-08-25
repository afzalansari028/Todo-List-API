package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// DB set up
func SetupDB() *sql.DB {
	db, err := sql.Open("mysql", "root:king@tcp(127.0.0.1:3306)/learn")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
