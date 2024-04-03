package http

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/amirhnajafiz/ghoster/internal/cexe"
	"github.com/amirhnajafiz/ghoster/internal/metrics"
	"golang.org/x/sync/semaphore"

	"github.com/gorilla/mux"
)

type Handler struct {
	Metrics         metrics.Metrics
	Semaphore       *semaphore.Weighted
	FunctionsDir    string
	DescriptionFile string
}

func (h Handler) ListFunctions(w http.ResponseWriter, r *http.Request) {
	functions, err := listDirectoryItems(h.FunctionsDir)
	if err != nil {
		h.error(w, http.StatusBadRequest, err)
		return
	}

	bytes, err := json.Marshal(functions)
	if err != nil {
		h.error(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}

func (h Handler) GetFunctionMarkdown(w http.ResponseWriter, r *http.Request) {
	// get param variables
	vars := mux.Vars(r)
	functionName := vars["function"]
	path := fmt.Sprintf("%s/%s/%s", h.FunctionsDir, functionName, h.DescriptionFile)

	if flag, err := fileOrDirExists(path); err != nil || !flag {
		h.error(w, http.StatusNotFound, nil)
		return
	}

	http.ServeFile(w, r, path)
}

func (h Handler) ExecuteFunction(w http.ResponseWriter, r *http.Request) {
	// get param variables
	vars := mux.Vars(r)
	functionName := vars["function"]
	path := fmt.Sprintf("%s/%s", h.FunctionsDir, functionName)

	if flag, err := fileOrDirExists(path); err != nil || !flag {
		h.error(w, http.StatusNotFound, nil)
		return
	}

	// parse json body
	decoder := json.NewDecoder(r.Body)
	var req Request

	if err := decoder.Decode(&req); err != nil {
		h.error(w, http.StatusBadRequest, err)
		return
	}

	args := []string{"run", "main.go"}
	args = append(args, req.Args...)

	// get a resource to continue
	ctx := context.Background()
	h.Semaphore.Acquire(ctx, 1)
	defer func() {
		h.Semaphore.Release(1)
	}()

	h.Metrics.AddFunctionCount(functionName, false)

	// function execute command and
	// get the command output
	bytes, duration, err := cexe.Execute(path, args)
	if err != nil {
		h.Metrics.AddFunctionCount(functionName, true)
		h.error(w, http.StatusBadGateway, err)
		return
	}

	h.Metrics.AddFunctionResponseTime(functionName, duration)

	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}

func (h Handler) Health(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (h Handler) error(w http.ResponseWriter, status int, err error) {
	if err != nil {
		log.Printf("ghoster handler error: %v\n", err)
	}

	w.WriteHeader(status)
}
