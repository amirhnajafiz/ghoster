package store

import (
	"encoding/json"
	"math/rand"
	"net/http"
	S "restapi/restapi/config/server"
	B "restapi/restapi/internal/models/book"
	"strconv"
)

// CreateBook : Create a New book
func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var book B.Book
	_ = json.NewDecoder(r.Body).Decode(&book)

	book.ID = strconv.Itoa(rand.Intn(10000000000)) // Mock id (not safe)
	S.Books = append(S.Books, book)

	_ = json.NewEncoder(w).Encode(book)
}
