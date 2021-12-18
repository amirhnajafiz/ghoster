package index

import (
	"encoding/json"
	"net/http"

	"restapi/restapi/config/server"
)

// GetBooks : Get all books
func GetBooks(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(server.Books)
}
