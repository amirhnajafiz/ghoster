package database

import (
	A "restapi/restapi/internal/models/author"
	B "restapi/restapi/internal/models/book"
)

// Mock : this will create a base slice of our books and authors
func Mock() []B.Book {
	var books []B.Book

	// Mock data
	books = append(books, B.Book{
		ID:    "1",
		Isbn:  "44502",
		Title: "Book One",
		Author: &A.Author{
			Firstname: "John",
			Lastname:  "Doe",
		},
	})
	books = append(books, B.Book{
		ID:    "2",
		Isbn:  "88727",
		Title: "Book Two",
		Author: &A.Author{
			Firstname: "Steve",
			Lastname:  "Smith",
		},
	})

	return books
}
