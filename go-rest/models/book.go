package models

import (
	"fmt"
	"koffe0522/go-projects/go-rest/database"
)

// Book は本データの構造体
type Book struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

// All is return Books
func All() []Book {
	books := []Book{}
	db := database.Db.Connect()
	db.Find(&books)
	if len(books) == 0 {
		fmt.Printf("not found users")
	}
	return books
}

// FindByID is return Book
func FindByID(id int) Book {
	book := Book{}
	db := database.Db.Connect()
	db.Where("id = ?", id).Find(&book)

	return book
}

// Create is create book
func Create(book Book) bool {
	db := database.Db.Connect()
	if err := db.Create(&book).Error; err != nil {
		return false
	}

	return true
}
