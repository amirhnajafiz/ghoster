// CMD:
//	This package sets up the main server of the application.
//	We load the configurations, and then we start the server.
//
///**
package main

import (
	"log"

	"restapi/restapi/config"
)

func main() {
	app := config.Config() // Getting the application server from config package

	// Starting the server
	log.Println("Server started ...")
	log.Fatal(app.ListenAndServe())
}
