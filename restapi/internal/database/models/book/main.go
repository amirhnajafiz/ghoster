package book

import (
	"gorm.io/gorm"
	"restapi/restapi/config/database"
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
	database.DB.First(Book{}, ID).Scan(&tempBook)
	return tempBook
}

func Add(tempBook Book) Book {
	database.DB.Create(&tempBook)
	return tempBook
}

func Put(tempBook Book, ID int) Book {
	book := Get(ID)
	database.DB.Model(&book).Updates(tempBook)
	return book
}

func Del(ID int) {
	database.DB.Delete(&Book{}, ID)
}
