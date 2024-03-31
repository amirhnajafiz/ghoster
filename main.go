package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/amirhnajafiz/ghoster/internal/config"
	"github.com/amirhnajafiz/ghoster/internal/file"
	internalHttp "github.com/amirhnajafiz/ghoster/internal/http"
	"github.com/amirhnajafiz/ghoster/internal/http/middleware"
	"github.com/amirhnajafiz/ghoster/internal/metrics"
	"github.com/amirhnajafiz/ghoster/internal/worker"

	"github.com/gorilla/mux"
)

func main() {
	// load env variables
	cfg := config.Load()

	// create a new mux router
	router := mux.NewRouter()

	// create an instance of internal handler
	h := internalHttp.Handler{
		Metrics: metrics.Register(cfg.MetricsNamespace, cfg.MetricsSubSystem),
		Pool:    worker.NewPool(cfg.PoolSize),
	}

	router.Use(middleware.Logging)
	router.Use(middleware.Metrics(h.Metrics))

	router.HandleFunc("/healthz", h.Health).Methods(http.MethodGet)
	router.HandleFunc("/functions", h.ListFunctions).Methods(http.MethodGet)
	router.HandleFunc("/functions/{function}", h.GetFunctionMarkdown).Methods(http.MethodGet)
	router.HandleFunc("/functions/{function}", h.ExecuteFunction).Methods(http.MethodPost)

	// create a new server
	srv := &http.Server{
		Handler: router,
		Addr:    fmt.Sprintf("127.0.0.1:%d", cfg.HTTPPort),
	}

	// register file server
	file.NewServer(cfg.FileServerPort)

	// register metrics server
	metrics.NewServer(cfg.MetricsPort)

	log.Printf("ghoster server started on %d ...\n", cfg.HTTPPort)

	// start the http server
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
