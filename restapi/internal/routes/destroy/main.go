package destroy

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	S "restapi/restapi/config/server"
)

// DeleteBook : Removing a book
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range S.Books {
		if item.ID == params["id"] {
			S.Books = append(S.Books[:index], S.Books[index+1:]...)
			break
		}
	}

	_ = json.NewEncoder(w).Encode(S.Books)
}
