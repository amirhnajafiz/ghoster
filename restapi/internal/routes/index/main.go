package index

import (
	"encoding/json"
	"net/http"
	"restapi/restapi/internal/database/models/book/methods"
)

// GetBooks : Get all books
func GetBooks(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(methods.All())
}
