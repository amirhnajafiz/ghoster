package config

import (
	"net/http"

	"restapi/restapi/config/database"
	"restapi/restapi/config/router"
	"restapi/restapi/config/server"
)

func Config() *http.Server {
	r := router.GetRouter()
	server.Books = database.Mock()
	app := server.GetServer(r)
	return app
}
