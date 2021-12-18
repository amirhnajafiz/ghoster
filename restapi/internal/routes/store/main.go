package store

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	"restapi/restapi/config/server"
	"restapi/restapi/internal/models/book"
)

// CreateBook : Create a New book
func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var tempBook book.Book
	_ = json.NewDecoder(r.Body).Decode(&tempBook)

	tempBook.ID = strconv.Itoa(rand.Intn(10000000000)) // Mock id (not safe)
	server.Books = append(server.Books, tempBook)

	_ = json.NewEncoder(w).Encode(tempBook)
}
