package main

import (
	"fmt"
	"go-personal-page/internals/config"
	"go-personal-page/internals/models"
	"go-personal-page/internals/routes"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found")
	}

	db, err := config.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	err = db.AutoMigrate(&models.Todo{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	e := echo.New()

	routes.SetupRoute(e)

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "3000"
	}

	serverAddress := fmt.Sprintf(":%s", port)
	log.Printf("Server started on %s", serverAddress)
	e.Start(serverAddress)
}
