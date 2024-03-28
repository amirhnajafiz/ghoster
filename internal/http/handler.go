package http

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/amirhnajafiz/ghoster/internal/metrics"
)

type Handler struct {
	Metrics metrics.Metrics
}

const functionsDir = "functions"

func (h Handler) ListFunctions(w http.ResponseWriter, r *http.Request) {
	functions, err := listDirectoryItems(functionsDir)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)

		log.Println(err)

		return
	}

	bytes, err := json.Marshal(functions)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		log.Println(err)

		return
	}

	h.Metrics.ListRequests.Add(1)

	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}

func (h Handler) Health(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}
