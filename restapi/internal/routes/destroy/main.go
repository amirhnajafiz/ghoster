package destroy

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"restapi/restapi/internal/database/models/book/methods"
	"strconv"
)

// DeleteBook : Removing a book
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	ID, _ := strconv.Atoi(params["id"])
	methods.Del(ID)

	_ = json.NewEncoder(w).Encode(methods.All())
}
