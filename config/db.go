package config

import (
	"database/sql"
	"fmt"
	"os"

	// load postgres driver
	_ "github.com/lib/pq"
)

var DB *sql.DB

type Config struct {
	DB_Username string
	DB_Password string
	DB_Port     string
	DB_Host     string
	DB_Name     string
}

func CreateConnection() *sql.DB {
	config := Config{
		DB_Username: os.Getenv("DB_USERNAME"),
		DB_Password: os.Getenv("DB_PASSWORD"),
		DB_Port:     os.Getenv("DB_PORT"),
		DB_Host:     os.Getenv("DB_HOST"),
		DB_Name:     os.Getenv("DB_NAME"),
	}

	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.DB_Host,
		config.DB_Port,
		config.DB_Username,
		config.DB_Password,
		config.DB_Name,
	)

	var err error
	// Open the connection
	DB, err = sql.Open("postgres", connectionString)

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
