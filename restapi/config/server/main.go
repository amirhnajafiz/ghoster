package server

import (
	"net/http"
	"restapi/restapi/config/router"
	B "restapi/restapi/internal/models/book"
)

var Books []B.Book

// GetServer : returns the server of the application
func GetServer() *http.Server {
	r := router.GetRouter()

	server := http.Server{
		Addr:    ":8000",
		Handler: r,
	}

	return &server
}
