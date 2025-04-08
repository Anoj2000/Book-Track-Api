package services

import (
	"book-api/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to test database")
	}
	db.AutoMigrate(&models.Book{})
	return db
}

func TestBookService_CreateBook(t *testing.T) {
	db := setupTestDB()
	service := NewBookService(db)

	book := models.Book{
		Title:  "Test Book",
		Author: "Test Author",
		Year:   2023,
	}

	err := service.CreateBook(&book)
	assert.NoError(t, err)
	assert.NotZero(t, book.ID)
}

func TestBookService_GetBookByID(t *testing.T) {
	db := setupTestDB()
	service := NewBookService(db)

	// Create a test book
	book := models.Book{
		Title:  "Test Book",
		Author: "Test Author",
		Year:   2023,
	}
	db.Create(&book)

	// Test getting the book
	retrievedBook, err := service.GetBookByID(book.ID)
	assert.NoError(t, err)
	assert.Equal(t, book.Title, retrievedBook.Title)
	assert.Equal(t, book.Author, retrievedBook.Author)
	assert.Equal(t, book.Year, retrievedBook.Year)
}

func TestBookService_GetAllBooks(t *testing.T) {
	db := setupTestDB()
	service := NewBookService(db)

	// Create test books
	books := []models.Book{
		{Title: "Book 1", Author: "Author 1", Year: 2000},
		{Title: "Book 2", Author: "Author 2", Year: 2001},
	}
	for _, book := range books {
		db.Create(&book)
	}

	// Test getting all books
	retrievedBooks, err := service.GetAllBooks()
	assert.NoError(t, err)
	assert.Len(t, retrievedBooks, 2)
}

func TestBookService_UpdateBook(t *testing.T) {
	db := setupTestDB()
	service := NewBookService(db)

	// Create a test book
	book := models.Book{
		Title:  "Original Title",
		Author: "Original Author",
		Year:   2000,
	}
	db.Create(&book)

	// Update the book
	updatedBook := models.Book{
		Title:  "Updated Title",
		Author: "Updated Author",
		Year:   2001,
	}
	err := service.UpdateBook(book.ID, updatedBook)
	assert.NoError(t, err)

	// Verify the update
	updated, err := service.GetBookByID(book.ID)
	assert.NoError(t, err)
	assert.Equal(t, "Updated Title", updated.Title)
	assert.Equal(t, "Updated Author", updated.Author)
	assert.Equal(t, 2001, updated.Year)
}

func TestBookService_DeleteBook(t *testing.T) {
	db := setupTestDB()
	service := NewBookService(db)

	// Create a test book
	book := models.Book{
		Title:  "Book to Delete",
		Author: "Author",
		Year:   2000,
	}
	db.Create(&book)

	// Delete the book
	err := service.DeleteBook(book.ID)
	assert.NoError(t, err)

	// Verify deletion
	_, err = service.GetBookByID(book.ID)
	assert.Error(t, err)
}