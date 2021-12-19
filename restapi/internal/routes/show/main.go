package show

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"restapi/restapi/internal/database/models/book/methods"
	"strconv"
)

// GetBook : Get a book from books struct
func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r) /// Get params
	id, _ := strconv.Atoi(params["id"])
	tempBook := methods.Get(id)

	_ = json.NewEncoder(w).Encode(tempBook)
}
