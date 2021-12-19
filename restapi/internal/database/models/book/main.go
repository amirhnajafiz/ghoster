package book

import (
	"gorm.io/gorm"
	"restapi/restapi/config/database"
	"restapi/restapi/internal/models/book"
)

// Book struct (Model)
type Book struct {
	gorm.Model
	Isbn   string
	Title  string
	Author string
}

func All() []Book {
	var books []Book
	_ = database.DB.Find(&books)
	return books
}

func Get(ID int) Book {
	tempBook := Book{}
	database.DB.First(book.Book{}, ID).Scan(&tempBook)
	return tempBook
}
