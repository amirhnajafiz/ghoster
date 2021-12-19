package update

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"restapi/restapi/internal/database/models/book"
	"restapi/restapi/internal/database/models/book/methods"
	"strconv"
)

// UpdateBook : updates a book in our slice
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	tempBook := book.Book{}
	_ = json.NewDecoder(r.Body).Decode(&tempBook)
	ID, _ := strconv.Atoi(params["id"])

	_ = json.NewEncoder(w).Encode(methods.Put(tempBook, ID))
}
