package main

import (
	"log"
	"net/http"

	"github.com/amirhnajafiz/restful-go/internal/config"
	"github.com/gorilla/mux"
)

// GetServer : returns the server of the application
func getServer(r *mux.Router) *http.Server {
	server := http.Server{
		Addr:    ":8000",
		Handler: r,
	}

	return &server
}

func main() {
	app := config.Config() // Getting the application server from config package

	// Starting the server
	log.Println("Server started ...")
	log.Fatal(app.ListenAndServe())
}
