// Package config
// 	This package creates the routers and the handlers
//  of the project.
//	After that is mocks the initialized data.
//  Finally, it will create a server and returns it.
//
///**
package config

import (
	"net/http"
	"restapi/restapi/internal/seeder"

	"restapi/restapi/config/database"
	"restapi/restapi/config/db"
	"restapi/restapi/config/router"
	"restapi/restapi/config/server"
)

// Config : sets up the server, router and database.
func Config() *http.Server {
	r := router.GetRouter()
	server.Books = database.Mock()
	db.Connect()
	seeder.Seed()
	app := server.GetServer(r)
	return app
}
