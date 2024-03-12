package main

import (
	"backend-marketplace-openidea/config"
	"backend-marketplace-openidea/routes"
	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	db := config.CreateConnection()

	defer db.Close()
	e := echo.New()
	routes.Routes(e, db)

	e.Logger.Fatal(e.Start(":8000"))
}
