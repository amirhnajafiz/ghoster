package destroy

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"restapi/restapi/internal/database/models/book"
	"strconv"
)

// DeleteBook : Removing a book
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	ID, _ := strconv.Atoi(params["id"])
	book.Del(ID)

	_ = json.NewEncoder(w).Encode(book.All())
}
