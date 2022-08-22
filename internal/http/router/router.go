package router

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/amirhnajafiz/restful-go/internal/models"
	"github.com/gorilla/mux"
)

// getBooks : Get all books
func getBooks(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	_ = json.NewEncoder(w).Encode(models.GetAllBooks())
}

// getBook : Get a book from books struct
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r) /// Get params
	id, _ := strconv.Atoi(params["id"])
	tempBook := models.GetBook(id)

	_ = json.NewEncoder(w).Encode(tempBook)
}

// createBook : Create a New book
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var tempBook models.Book

	_ = json.NewDecoder(r.Body).Decode(&tempBook)

	_ = json.NewEncoder(w).Encode(models.AddBook(tempBook))
}

// updateBook : updates a book in our slice
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	tempBook := models.Book{}

	_ = json.NewDecoder(r.Body).Decode(&tempBook)
	ID, _ := strconv.Atoi(params["id"])

	_ = json.NewEncoder(w).Encode(models.PutBook(tempBook, ID))
}

// deleteBook : Removing a book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	ID, _ := strconv.Atoi(params["id"])

	models.DelBook(ID)

	_ = json.NewEncoder(w).Encode(models.GetAllBooks())
}

// GetRouter : returns the application router
func GetRouter() *mux.Router {
	// Init router
	r := mux.NewRouter()

	// Route handlers / Endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	return r
}
