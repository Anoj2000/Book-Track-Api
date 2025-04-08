package handlers

import (
	"book-api/models"
	"book-api/services"
	"errors"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type BookHandler struct {
	service *services.BookService
}

func NewBookHandler(service *services.BookService) *BookHandler {
	return &BookHandler{service: service}
}

// Enhanced error response structure
type ErrorResponse struct {
	Error   string      `json:"error"`
	Details interface{} `json:"details,omitempty"`
}

func (h *BookHandler) CreateBook(c *fiber.Ctx) error {
	var book models.Book
	
	if err := c.BodyParser(&book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Error:   "Invalid request body",
			Details: "Expected JSON format: {\"title\":\"string\",\"author\":\"string\",\"year\":number}",
		})
	}

	// Validate required fields
	if strings.TrimSpace(book.Title) == "" || strings.TrimSpace(book.Author) == "" {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Error: "Title and author are required fields",
		})
	}

	if err := h.service.CreateBook(&book); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{
			Error: "Failed to create book",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(book)
}

func (h *BookHandler) GetAllBooks(c *fiber.Ctx) error {
	books, err := h.service.GetAllBooks()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{
			Error: "Failed to retrieve books",
		})
	}
	
	if len(books) == 0 {
		return c.Status(fiber.StatusOK).JSON([]interface{}{})
	}
	
	return c.Status(fiber.StatusOK).JSON(books)
}

func (h *BookHandler) GetBookByID(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil || id == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Error:   "Invalid book ID",
			Details: "ID must be a positive integer",
		})
	}

	book, err := h.service.GetBookByID(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(ErrorResponse{
				Error: "Book not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{
			Error: "Failed to retrieve book",
		})
	}

	return c.Status(fiber.StatusOK).JSON(book)
}

func (h *BookHandler) UpdateBook(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil || id == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Error:   "Invalid book ID",
			Details: "ID must be a positive integer",
		})
	}

	var updatedBook models.Book
	if err := c.BodyParser(&updatedBook); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Error:   "Invalid request body",
			Details: "Expected JSON format with update fields",
		})
	}

	if err := h.service.UpdateBook(uint(id), updatedBook); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(ErrorResponse{
				Error: "Book not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{
			Error: "Failed to update book",
		})
	}

	return c.SendStatus(fiber.StatusOK)
}

func (h *BookHandler) DeleteBook(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil || id == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Error:   "Invalid book ID",
			Details: "ID must be a positive integer",
		})
	}

	if err := h.service.DeleteBook(uint(id)); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(ErrorResponse{
				Error: "Book not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{
			Error: "Failed to delete book",
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (h *BookHandler) GetBooksPaginated(c *fiber.Ctx) error {
	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil || page < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Error:   "Invalid page number",
			Details: "Page must be a positive integer",
		})
	}

	pageSize, err := strconv.Atoi(c.Query("pageSize", "10"))
	if err != nil || pageSize < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Error:   "Invalid page size",
			Details: "PageSize must be a positive integer",
		})
	}

	// Limit maximum page size
	if pageSize > 100 {
		pageSize = 100
	}

	books, total, err := h.service.GetBooksPaginated(page, pageSize)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{
			Error: "Failed to fetch paginated books",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":       books,
		"total":      total,
		"page":       page,
		"page_size":  pageSize,
		"total_pages": (int(total) + pageSize - 1) / pageSize,
	})
}

func (h *BookHandler) SearchBooks(c *fiber.Ctx) error {
	query := strings.TrimSpace(c.Query("q"))
	if query == "" {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Error: "Search query cannot be empty",
		})
	}

	// Prevent overly broad searches
	if len(query) < 2 {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Error: "Search query must be at least 2 characters long",
		})
	}

	books, err := h.service.SearchBooks(query)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{
			Error: "Failed to search books",
		})
	}

	return c.Status(fiber.StatusOK).JSON(books)
}

func (h *BookHandler) DeleteAllBooks(c *fiber.Ctx) error {
	// Add security check in production
	// if !isAdmin(c) { return c.Status(403)... }
	
	if err := h.service.DeleteAllBooks(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{
			Error: "Failed to delete all books",
		})
	}
	
	return c.SendStatus(fiber.StatusNoContent)
}