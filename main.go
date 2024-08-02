package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ealekseychik/mnemosyne/internal/handlers"
	"github.com/ealekseychik/mnemosyne/internal/models"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	dbHost := os.Getenv("POSTGRES_HOST")
	dbPort := os.Getenv("POSTGRES_PORT")

	// Wait for the database to start
	// TODO: Switch to wait-for-it.sh
	time.Sleep(5 * time.Second)

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		dbHost, dbUser, dbPassword, dbName, dbPort,
	)

	models.InitDB(dsn)
	defer models.CloseDB()

	models.SeedDB()

	handlers.StartPeriodicTasks()

	router := handlers.SetupRouter()

	router.Run(":8080")
}
