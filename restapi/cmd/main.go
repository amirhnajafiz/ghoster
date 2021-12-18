package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	A "restapi/restapi/internal/models/author"
	B "restapi/restapi/internal/models/book"
	"restapi/restapi/internal/routes/index"
	"restapi/restapi/internal/routes/show"
	"restapi/restapi/internal/routes/store"
	"restapi/restapi/internal/routes/update"
)

// Init books var as a slice Book struct
var books []B.Book

// Removing a book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}

	_ = json.NewEncoder(w).Encode(books)
}

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
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Println("Server started ...")
	// Starting the server
	log.Fatal(http.ListenAndServe(":8000", r))
}
