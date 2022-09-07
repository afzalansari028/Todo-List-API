package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "king"
	dbname   = "todo-appdb"
)

// DB set up
func SetupDB() *sql.DB {
	//connection
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	//open database
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		log.Fatal(err)
	}
	// // close database
	// defer db.Close()
	// verify db
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("db connected...")
	}
	return db
}
