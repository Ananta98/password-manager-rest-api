package main

import (
	"log"
	"password-manager/config"
	"password-manager/routes"

	"github.com/joho/godotenv"
	"github.com/swaggo/swag/example/override/docs"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	docs.SwaggerInfo.Title = "Swagger REST API Password Manager"
	docs.SwaggerInfo.Description = "This is Golang Backend Password Manager."
	docs.SwaggerInfo.Version = "1.0"

	db := config.ConnectDatabase()
	sqlDb, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDb.Close()
	r := routes.SetupRoutes(db)
	r.Run()
}
