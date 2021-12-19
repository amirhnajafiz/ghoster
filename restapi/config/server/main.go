package server

import (
	"github.com/gorilla/mux"
	"net/http"
)

// GetServer : returns the server of the application
func GetServer(r *mux.Router) *http.Server {
	server := http.Server{
		Addr:    ":8000",
		Handler: r,
	}
	return &server
}
