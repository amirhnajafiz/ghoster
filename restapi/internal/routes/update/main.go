package update

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	S "restapi/restapi/config/server"
	B "restapi/restapi/internal/models/book"
)

// UpdateBook : updates a book in our slice
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range S.Books {
		if item.ID == params["id"] {
			S.Books = append(S.Books[:index], S.Books[index+1:]...)
			var book B.Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			S.Books = append(S.Books, book)

			_ = json.NewEncoder(w).Encode(book)
			return
		}
	}

	_ = json.NewEncoder(w).Encode(S.Books)
}
