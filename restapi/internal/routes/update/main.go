package update

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"

	"restapi/restapi/config/server"
	"restapi/restapi/internal/models/book"
)

// UpdateBook : updates a book in our slice
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range server.Books {
		if item.ID == params["id"] {
			server.Books = append(server.Books[:index], server.Books[index+1:]...)

			var tempBook book.Book
			_ = json.NewDecoder(r.Body).Decode(&tempBook)
			tempBook.ID = params["id"]
			server.Books = append(server.Books, tempBook)

			_ = json.NewEncoder(w).Encode(tempBook)
			return
		}
	}

	_ = json.NewEncoder(w).Encode(server.Books)
}
