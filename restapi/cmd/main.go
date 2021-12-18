package main

import (
	"log"
	"restapi/restapi/config/server"

	A "restapi/restapi/internal/models/author"
	B "restapi/restapi/internal/models/book"
)

// Init books var as a slice Book struct
var books []B.Book

func main() {
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

	app := server.GetServer()
	log.Println("Server started ...")
	// Starting the server
	log.Fatal(app.ListenAndServe())
}
