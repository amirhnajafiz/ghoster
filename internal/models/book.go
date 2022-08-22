package models

import "github.com/amirhnajafiz/restful-go/internal/database"

// Book struct (Model)
type Book struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Isbn     string `json:"isbn"`
	Title    string `json:"title"`
	AuthorID int    `json:"author_id"`
	Author   Author `gorm:"references:ID"`
}

func GetAllBooks() []Book {
	var books []Book

	_ = database.DB.Find(&books)

	return books
}

func GetBook(ID int) Book {
	var tempBook Book

	database.DB.First(&tempBook, ID)

	return tempBook
}

func AddBook(tempBook Book) Book {
	database.DB.Create(&tempBook)

	return tempBook
}

func PutBook(tempBook Book, ID int) Book {
	temp := GetBook(ID)

	database.DB.Model(&temp).Updates(tempBook)

	return temp
}

func DelBook(ID int) {
	database.DB.Delete(&Book{}, ID)
}
