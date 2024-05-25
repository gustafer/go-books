package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gustafer/go-books/cmd/api-server/handlers"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, db *gorm.DB) {
	r := handlers.Repository{
		DB: db,
	}

	api := app.Group("/api")

	api.Post("/book", r.CreateBook)
	api.Delete("/book/:id", r.DeleteBook)
	api.Get("/book/:id", r.GetBookByID)
	api.Get("/books", r.GetBooks)
}
