package database

import (
	"book-api/models"

	"github.com/glebarez/sqlite" // Pure Go SQLite
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() error {
    var err error
    DB, err = gorm.Open(sqlite.Open("books.db"), &gorm.Config{})
    if err != nil {
        return err
    }
    return DB.AutoMigrate(&models.Book{})
}