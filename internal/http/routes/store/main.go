package store

import (
	"encoding/json"
	"net/http"
	"restapi/restapi/internal/database/models/book"
	"restapi/restapi/internal/database/models/book/methods"
)

// CreateBook : Create a New book
func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var tempBook book.Book
	_ = json.NewDecoder(r.Body).Decode(&tempBook)

	_ = json.NewEncoder(w).Encode(methods.Add(tempBook))
}
