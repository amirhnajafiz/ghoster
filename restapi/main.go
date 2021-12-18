package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	_ "math/rand"
	"net/http"
	_ "strconv"
)

// Book struct (Model)
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Author struct
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// Init books var as a slice Book struct
var books []Book

// Get all books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	log.Fatal(json.NewEncoder(w).Encode(books))
}

func getBook(w http.ResponseWriter, r *http.Request) {

}

func createBook(w http.ResponseWriter, r *http.Request) {

}

func updateBook(w http.ResponseWriter, r *http.Request) {

}

func deleteBook(w http.ResponseWriter, r *http.Request) {

}

func main() {
	// Init router
	r := mux.NewRouter()

	// Mock data
	books = append(books, Book{
		ID:    "1",
		Isbn:  "44502",
		Title: "Book One",
		Author: &Author{
			Firstname: "John",
			Lastname:  "Doe",
		},
	})
	books = append(books, Book{
		ID:    "2",
		Isbn:  "88727",
		Title: "Book Two",
		Author: &Author{
			Firstname: "Steve",
			Lastname:  "Smith",
		},
	})

	// Route handlers / Endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	// Starting the server
	log.Fatal(http.ListenAndServe(":8000", r))
}
