package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" // Driver for MySQL
)

const (
	dbHost = "tcp(127.0.0.1:3306)" // localhost and port (Default MySQL: 3306)
	dbName = "golang"              // Must have created db in MySQl
	dbUser = "username"            // Enter your username for MySQL
	dbPass = "password"            // Enter your password
)

func main() {
	dsn := dbUser + ":" + dbPass + "@" + dbHost + "/" + dbName + "?charset=utf8"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Success! Connected to database...")
}
