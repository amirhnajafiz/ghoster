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

func Show(ID int) Book {
	tempBook := Book{}
	database.DB.First(book.Book{}, ID).Scan(&tempBook)
	return tempBook
}
