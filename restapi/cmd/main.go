package main

import (
	"log"
	"restapi/restapi/config"
)

func main() {
	app := config.Config()
	log.Println("Server started ...")
	// Starting the server
	log.Fatal(app.ListenAndServe())
}
