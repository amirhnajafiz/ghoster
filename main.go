package main

import (
	"log"
	"net/http"

	internalHttp "github.com/amirhnajafiz/ghoster/internal/http"

	"github.com/gorilla/mux"
)

func main() {
	// create a new mux router
	router := mux.NewRouter()

	// create an instance of internal handler
	h := internalHttp.Handler{}

	router.Methods(http.MethodGet).HandlerFunc(h.ListFunctions)

	// create a new server
	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:5000",
	}

	// start the http server
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
