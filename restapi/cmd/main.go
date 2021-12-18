package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"

	A "restapi/restapi/internal/models/author"
	B "restapi/restapi/internal/models/book"
	"restapi/restapi/internal/routes/destroy"
	"restapi/restapi/internal/routes/index"
	"restapi/restapi/internal/routes/show"
	"restapi/restapi/internal/routes/store"
	"restapi/restapi/internal/routes/update"
)

// Init books var as a slice Book struct
var books []B.Book

func main() {
	// Init router
	r := mux.NewRouter()

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

	// Route handlers / Endpoints
	r.HandleFunc("/api/books", index.GetBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", show.GetBook).Methods("GET")
	r.HandleFunc("/api/books", store.CreateBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", update.UpdateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", destroy.DeleteBook).Methods("DELETE")

	log.Println("Server started ...")
	// Starting the server
	log.Fatal(http.ListenAndServe(":8000", r))
}
