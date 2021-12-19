package methods

import (
	"restapi/restapi/config/database"
	"restapi/restapi/internal/database/models/book"
)

func All() []book.Book {
	var books []book.Book
	_ = database.DB.Find(&books)
	return books
}

func Get(ID int) book.Book {
	var tempBook book.Book
	database.DB.First(&tempBook, ID)
	return tempBook
}

func Add(tempBook book.Book) book.Book {
	database.DB.Create(&tempBook)
	return tempBook
}

func Put(tempBook book.Book, ID int) book.Book {
	temp := Get(ID)
	database.DB.Model(&temp).Updates(tempBook)
	return temp
}

func Del(ID int) {
	database.DB.Delete(&book.Book{}, ID)
}
