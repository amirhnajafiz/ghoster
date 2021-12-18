package main

import (
	_ "encoding/json"
	"github.com/gorilla/mux"
	_ "log"
	_ "math/rand"
	_ "net/http"
	_ "strconv"
)

func main() {
	// Init router
	r := mux.NewRouter()

	// Route handlers / Endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")
}
