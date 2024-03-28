package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/amirhnajafiz/ghoster/internal/config"
	internalHttp "github.com/amirhnajafiz/ghoster/internal/http"

	"github.com/gorilla/mux"
)

func main() {
	// load env variables
	cfg := config.Load()

	// create a new mux router
	router := mux.NewRouter()

	// create an instance of internal handler
	h := internalHttp.Handler{}

	router.Methods(http.MethodGet).HandlerFunc(h.ListFunctions)

	// create a new server
	srv := &http.Server{
		Handler: router,
		Addr:    fmt.Sprintf("127.0.0.1:%d", cfg.HTTPPort),
	}

	// start the http server
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
