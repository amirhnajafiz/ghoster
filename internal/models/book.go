package models

import "gorm.io/gorm"

// Book struct (Model)
type Book struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Isbn     string `json:"isbn"`
	Title    string `json:"title"`
	AuthorID int    `json:"author_id"`
	Author   Author `gorm:"references:ID"`
}

func GetAllBooks(db *gorm.DB) []Book {
	var books []Book

	_ = db.Find(&books)

	return books
}

func GetBook(ID int, db *gorm.DB) Book {
	var tempBook Book

	db.First(&tempBook, ID)

	return tempBook
}

func AddBook(tempBook Book, db *gorm.DB) Book {
	db.Create(&tempBook)

	return tempBook
}

func PutBook(tempBook Book, ID int, db *gorm.DB) Book {
	temp := GetBook(ID, db)

	db.Model(&temp).Updates(tempBook)

	return temp
}

func DelBook(ID int, db *gorm.DB) {
	db.Delete(&Book{}, ID)
}
