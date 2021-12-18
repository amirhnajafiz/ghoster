package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	A "restapi/restapi/internal/models/author"
	B "restapi/restapi/internal/models/book"
)

// Init books var as a slice Book struct
var books []B.Book

// Get all books
func getBooks(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(books)
}

// Get a book from books struct
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r) /// Get params

	// Loop in books and file with id
	for _, item := range books {
		if item.ID == params["id"] {
			_ = json.NewEncoder(w).Encode(item)
			return
		}
	}

	_ = json.NewEncoder(w).Encode(&B.Book{})
}

// Create a New book
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var book B.Book
	_ = json.NewDecoder(r.Body).Decode(&book)

	book.ID = strconv.Itoa(rand.Intn(10000000000)) // Mock id (not safe)
	books = append(books, book)

	_ = json.NewEncoder(w).Encode(book)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book B.Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book)

			_ = json.NewEncoder(w).Encode(book)
			return
		}
	}

	_ = json.NewEncoder(w).Encode(books)
}

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
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Println("Server started ...")
	// Starting the server
	log.Fatal(http.ListenAndServe(":8000", r))
}
