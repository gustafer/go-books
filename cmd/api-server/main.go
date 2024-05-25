package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gustafer/go-books/cmd/api-server/models"
	"github.com/gustafer/go-books/cmd/api-server/routes"
	"github.com/gustafer/go-books/cmd/api-server/storage"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	config := &storage.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SLLMODE"),
	}

	db, err := storage.NewConnection(config)

	if err != nil {
		log.Fatalf("Could not connect to db, \n Reason: %v", err)
	}

	if err = models.MigrateBooks(db); err != nil {
		log.Fatalf("Could not migrate db, \n Reason: %v", err)
	}

	app := fiber.New()
	routes.SetupRoutes(app, db)
	app.Listen(":8080")

}
