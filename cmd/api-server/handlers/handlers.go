package handlers

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gustafer/go-books/cmd/api-server/models"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

type Book struct {
	Author    string `json:"author"`
	Title     string `json:"title"`
	Publisher string `json:"publisher"`
}

func (r *Repository) CreateBook(context *fiber.Ctx) error {
	book := Book{}

	err := context.BodyParser(&book)

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "request failed"})
		return err
	}

	err = r.DB.Create(&book).Error

	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "could not create book"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "book has been added."})
	return nil
}
func (r *Repository) DeleteBook(context *fiber.Ctx) error {
	bookModel := models.Books{}
	id := context.Params("id")

	if id == "" {
		context.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"message": "no id was provided",
		})
		return nil
	}

	err := r.DB.Delete(bookModel, id).Error

	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": fmt.Sprintf("could not delete the book with id: %v", id)})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "book deleted with ease"})

	return nil
}

func (r *Repository) GetBooks(context *fiber.Ctx) error {
	bookModels := &[]models.Books{}

	err := r.DB.Find(bookModels).Error

	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "could not get the books"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "books fetched succesfully",
		"data":    bookModels,
	})
	return nil
}

func (r *Repository) GetBookByID(context *fiber.Ctx) error {
	bookModel := &[]models.Books{}
	id := context.Params("id")

	if id == "" {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "no id was provided"})
	}

	err := r.DB.Where("id = ?", id).First(bookModel).Error

	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "could not get the book"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "books found by id succesfully",
		"data":    bookModel,
	})
	return nil
}
