package database

import (
	"restapi/restapi/internal/models/author"
	"restapi/restapi/internal/models/book"
)

// Mock : this will create a base slice of our books and authors
func Mock() []book.Book {
	var books []book.Book

	// Mock data
	books = append(books, book.Book{
		ID:    "1",
		Isbn:  "44502",
		Title: "Book One",
		Author: &author.Author{
			Firstname: "John",
			Lastname:  "Doe",
		},
	})
	books = append(books, book.Book{
		ID:    "2",
		Isbn:  "88727",
		Title: "Book Two",
		Author: &author.Author{
			Firstname: "Steve",
			Lastname:  "Smith",
		},
	})

	return books
}
