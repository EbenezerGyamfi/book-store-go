package models

import (
	"fmt"

	"example.com/bookstore/pkg/config"
	"gorm.io/gorm"
)


var db *gorm.DB

type Book struct {
	gorm.Model


	ID uint `gorm:"primary_key"`
	Title string
	Author string
	Price float64
	
}

func init()  {
	config.Connect()
	config.Db.AutoMigrate(&Book{})
}


func (b *Book) Save()  *Book{



    // Ensure the database connection is initialized
    if config.Db == nil {
        panic("Database connection is not initialized")
    }

    // Use the global database connection
    result := config.Db.Create(&b)

    // Check for errors during the query
    if result.Error != nil {
        panic("Failed to retrieve books: " + result.Error.Error())
    }

    return b
}	


func GetBooks() ([]Book, error) {
    var books []Book

    // Check if the database connection is initialized
    if config.Db == nil {
        return nil, fmt.Errorf("database connection is not initialized")
    }

    // Perform the query
    result := config.Db.Find(&books)
    if result.Error != nil {
        return nil, fmt.Errorf("failed to retrieve books: %w", result.Error)
    }

    return books, nil
}

func GetBookById(id int64) (*Book, error) {
    var getBook Book

    // Ensure the database connection is initialized
    if config.Db == nil {
        return nil, fmt.Errorf("database connection is not initialized")
    }

    // Query the book by ID
    result := config.Db.First(&getBook, id)
    if result.Error != nil {
        return nil, fmt.Errorf("failed to find book: %v", result.Error)
    }

    return &getBook, nil
}


func DeleteBook(id int64) Book {
    var book Book

    // Ensure the database connection is initialized
    if config.Db == nil {
        panic("Database connection is not initialized")
    }

    // Find the book to ensure it exists before deletion
    result := config.Db.First(&book, id)
    if result.Error != nil {
        panic("Book not found: " + result.Error.Error())
    }

    // Delete the book
    result = config.Db.Delete(&book)
    if result.Error != nil {
        panic("Failed to delete the book: " + result.Error.Error())
    }

    return book
}
func (b *Book) Update(id int64) *Book {
    // Ensure the database connection is initialized
    if config.Db == nil {
        panic("Database connection is not initialized")
    }

    // Find the existing record by ID
    var existingBook Book
    result := config.Db.First(&existingBook, id)
    if result.Error != nil {
        panic("Book not found: " + result.Error.Error())
    }

    // Update the existing record with new values
    result = config.Db.Model(&existingBook).Updates(b)
    if result.Error != nil {
        panic("Failed to update the book: " + result.Error.Error())
    }

    return &existingBook
}