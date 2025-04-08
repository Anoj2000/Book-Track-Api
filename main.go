package main

import (
	"book-api/database"
	"book-api/handlers"
	"book-api/services"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Add root route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Welcome to Book API",
			"routes": fiber.Map{
				"GET /api/books":       "List all books",
				"GET /api/books/:id":  "Get a book",
				"POST /api/books":     "Create a book",
				"PUT /api/books/:id":  "Update a book",
				"DELETE /api/books/:id": "Delete a book",
			},
		})
	})

	// Initialize database
	if err := database.Connect(); err != nil {
		panic("Failed to connect to database")
	}

	// Initialize services and handlers
	bookService := services.NewBookService(database.DB)
	bookHandler := handlers.NewBookHandler(bookService)

	// API routes
	api := app.Group("/api")
	api.Post("/books", bookHandler.CreateBook)
	api.Get("/books", bookHandler.GetAllBooks)
	api.Get("/books/:id", bookHandler.GetBookByID)
	api.Get("/books/paginated", bookHandler.GetBooksPaginated)
	api.Get("/books/search", bookHandler.SearchBooks)
	api.Put("/books/:id", bookHandler.UpdateBook)
	api.Delete("/books/:id", bookHandler.DeleteBook)
	api.Delete("/books", bookHandler.DeleteAllBooks) // Add this line

	// Start server
	app.Listen(":3000")
}