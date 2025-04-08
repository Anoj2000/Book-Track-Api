package services

import (
	"book-api/models"
	"gorm.io/gorm"
)

type BookService struct {
	db *gorm.DB
}

func NewBookService(db *gorm.DB) *BookService {
	return &BookService{db: db}
}

func (s *BookService) CreateBook(book *models.Book) error {
	return s.db.Create(book).Error
}

func (s *BookService) GetAllBooks() ([]models.Book, error) {
	var books []models.Book
	err := s.db.Find(&books).Error
	return books, err
}

func (s *BookService) GetBookByID(id uint) (models.Book, error) {
	var book models.Book
	err := s.db.First(&book, id).Error
	return book, err
}

func (s *BookService) UpdateBook(id uint, updatedBook models.Book) error {
	var book models.Book
	if err := s.db.First(&book, id).Error; err != nil {
		return err
	}
	return s.db.Model(&book).Updates(updatedBook).Error
}

func (s *BookService) DeleteBook(id uint) error {
	return s.db.Delete(&models.Book{}, id).Error
}