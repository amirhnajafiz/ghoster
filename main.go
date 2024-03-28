package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/amirhnajafiz/ghoster/internal/config"
	internalHttp "github.com/amirhnajafiz/ghoster/internal/http"
	"github.com/amirhnajafiz/ghoster/internal/metrics"

	"github.com/gorilla/mux"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.RequestURI != "/healthz" {
			log.Println(r.RequestURI)
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	// load env variables
	cfg := config.Load()

	// create a new mux router
	router := mux.NewRouter()

	// create an instance of internal handler
	h := internalHttp.Handler{
		Metrics: metrics.Register(cfg.MetricsNamespace, cfg.MetricsSubSystem),
	}

	router.Use(loggingMiddleware)

	router.HandleFunc("/healthz", h.Health).Methods(http.MethodGet)
	router.HandleFunc("/list", h.ListFunctions).Methods(http.MethodGet)

	// create a new server
	srv := &http.Server{
		Handler: router,
		Addr:    fmt.Sprintf("127.0.0.1:%d", cfg.HTTPPort),
	}

	// register metrics server
	metrics.NewServer(cfg.MetricsPort)

	log.Printf("ghoster server started on %d ...\n", cfg.HTTPPort)

	// start the http server
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
