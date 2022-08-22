package router

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/amirhnajafiz/restful-go/internal/models"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type router struct {
	db *gorm.DB
}

// getBooks : Get all books
func (router *router) getBooks(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	_ = json.NewEncoder(w).Encode(models.GetAllBooks(router.db))
}

// getBook : Get a book from books struct
func (router *router) getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r) /// Get params
	id, _ := strconv.Atoi(params["id"])
	tempBook := models.GetBook(id, router.db)

	_ = json.NewEncoder(w).Encode(tempBook)
}

// createBook : Create a New book
func (router *router) createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var tempBook models.Book

	_ = json.NewDecoder(r.Body).Decode(&tempBook)

	_ = json.NewEncoder(w).Encode(models.AddBook(tempBook, router.db))
}

// updateBook : updates a book in our slice
func (router *router) updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	tempBook := models.Book{}

	_ = json.NewDecoder(r.Body).Decode(&tempBook)
	ID, _ := strconv.Atoi(params["id"])

	_ = json.NewEncoder(w).Encode(models.PutBook(tempBook, ID, router.db))
}

// deleteBook : Removing a book
func (router *router) deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	ID, _ := strconv.Atoi(params["id"])

	models.DelBook(ID, router.db)

	_ = json.NewEncoder(w).Encode(models.GetAllBooks(router.db))
}

// GetRouter : returns the application router
func GetRouter(db *gorm.DB) *mux.Router {
	// Init router
	r := mux.NewRouter()

	router := &router{
		db: db,
	}

	// Route handlers / Endpoints
	r.HandleFunc("/api/books", router.getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", router.getBook).Methods("GET")
	r.HandleFunc("/api/books", router.createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", router.updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", router.deleteBook).Methods("DELETE")

	return r
}
