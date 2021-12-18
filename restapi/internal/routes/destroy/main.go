package destroy

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"

	"restapi/restapi/config/server"
)

// DeleteBook : Removing a book
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range server.Books {
		if item.ID == params["id"] {
			server.Books = append(server.Books[:index], server.Books[index+1:]...)
			break
		}
	}

	_ = json.NewEncoder(w).Encode(server.Books)
}
