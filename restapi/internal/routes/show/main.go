package show

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"

	"restapi/restapi/config/server"
	"restapi/restapi/internal/models/book"
)

// GetBook : Get a book from books struct
func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r) /// Get params

	// Loop in books and file with id
	for _, item := range server.Books {
		if item.ID == params["id"] {
			_ = json.NewEncoder(w).Encode(item)
			return
		}
	}

	_ = json.NewEncoder(w).Encode(&book.Book{})
}
