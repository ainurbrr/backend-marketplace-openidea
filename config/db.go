package config

import (
	"database/sql"
	"fmt"
	"os"

	// load postgres driver
	_ "github.com/lib/pq"
)

var DB *sql.DB

func CreateConnection() *sql.DB {
	var err error
	// Open the connection
	DB, err = sql.Open("postgres", os.Getenv("POSTGRES_URL"))

	if err != nil {
		panic(err)
	}

	// check the connection
	err = DB.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	// return the connection
	return DB
}
