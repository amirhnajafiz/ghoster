package main

import (
	"log"
	"restapi/restapi/config/server"
)

func main() {
	app := server.GetServer()
	log.Println("Server started ...")
	// Starting the server
	log.Fatal(app.ListenAndServe())
}
