package main

import (
	"log"
	"net/http"

	"github.com/amirhnajafiz/restful-go/internal/database"
	"github.com/amirhnajafiz/restful-go/internal/http/router"
)

// GetServer : returns the server of the application
func getServer() *http.Server {
	db := database.Connect(false)
	r := router.GetRouter(db)

	server := http.Server{
		Addr:    ":8000",
		Handler: r,
	}

	return &server
}

func main() {
	app := getServer()

	// Starting the server
	log.Println("Server started ...")
	log.Fatal(app.ListenAndServe())
}
