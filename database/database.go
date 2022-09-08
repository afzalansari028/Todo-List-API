package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "ec2-35-170-146-54.compute-1.amazonaws.com"
	port     = 5432
	user     = "ptsovhcdqfhwpw"
	password = "295c2f83522f0448515704e0f72615a92c1bd4cb57ee44ee5e785d0682ac57ad"
	dbname   = "dd8313ifetngdm"
)

// const (
// 	host     = "localhost"
// 	port     = 5432
// 	user     = "postgres"
// 	password = "king"
// 	dbname   = "todo-appdb"
// )

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
