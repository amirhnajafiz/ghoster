package router

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/amirhnajafiz/restful-go/internal/models"
	"github.com/gorilla/mux"
)

// GetBooks : Get all books
func GetBooks(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(models.GetAllBooks())
}

// GetBook : Get a book from books struct
func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r) /// Get params
	id, _ := strconv.Atoi(params["id"])
	tempBook := models.GetBook(id)

	_ = json.NewEncoder(w).Encode(tempBook)
}

// CreateBook : Create a New book
func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var tempBook models.Book
	_ = json.NewDecoder(r.Body).Decode(&tempBook)

	_ = json.NewEncoder(w).Encode(models.AddBook(tempBook))
}

// UpdateBook : updates a book in our slice
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	tempBook := models.Book{}
	_ = json.NewDecoder(r.Body).Decode(&tempBook)
	ID, _ := strconv.Atoi(params["id"])

	_ = json.NewEncoder(w).Encode(models.PutBook(tempBook, ID))
}

// DeleteBook : Removing a book
func DeleteBook(w http.ResponseWriter, r *http.Request) {
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
	r.HandleFunc("/api/books", GetBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", GetBook).Methods("GET")
	r.HandleFunc("/api/books", CreateBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", UpdateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", DeleteBook).Methods("DELETE")

	return r
}
