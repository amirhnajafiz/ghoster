package server

import (
	"github.com/gorilla/mux"
	"net/http"

	B "restapi/restapi/internal/models/book"
)

var Books []B.Book

// GetServer : returns the server of the application
func GetServer(r *mux.Router) *http.Server {
	server := http.Server{
		Addr:    ":8000",
		Handler: r,
	}
	return &server
}
