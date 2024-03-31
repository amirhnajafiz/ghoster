package http

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"time"

	"github.com/amirhnajafiz/ghoster/internal/metrics"
	"github.com/amirhnajafiz/ghoster/internal/worker"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/gorilla/mux"
)

type Handler struct {
	Metrics metrics.Metrics
	Pool    worker.Pool
}

const (
	functionsDir    = "functions"
	descriptionFile = "README.md"
)

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

	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}

func (h Handler) GetFunctionMarkdown(w http.ResponseWriter, r *http.Request) {
	// get param variables
	vars := mux.Vars(r)
	functionName := vars["function"]
	path := fmt.Sprintf("%s/%s/%s", functionsDir, functionName, descriptionFile)

	if flag, err := fileOrDirExists(path); err != nil || !flag {
		w.WriteHeader(http.StatusNotFound)

		return
	}

	http.ServeFile(w, r, path)
}

func (h Handler) ExecuteFunction(w http.ResponseWriter, r *http.Request) {
	// get param variables
	vars := mux.Vars(r)
	functionName := vars["function"]
	path := fmt.Sprintf("%s/%s", functionsDir, functionName)

	if flag, err := fileOrDirExists(path); err != nil || !flag {
		w.WriteHeader(http.StatusNotFound)

		return
	}

	// parse json body
	decoder := json.NewDecoder(r.Body)
	var req Request

	if err := decoder.Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)

		log.Println(err)

		return
	}

	args := []string{"run", "main.go"}
	args = append(args, req.Args...)

	// get a resource to continue
	h.Pool.Pull()
	defer h.Pool.Free()

	// function execute command
	cmd := exec.Command("go", args...)
	cmd.Dir = fmt.Sprintf("%s/%s", functionsDir, functionName)

	h.Metrics.FunctionCount.With(prometheus.Labels{"function": functionName}).Add(1)

	now := time.Now()

	// get the command output
	bytes, err := cmd.Output()
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)

		log.Println(err)

		h.Metrics.FunctionFailure.With(prometheus.Labels{"function": functionName}).Add(1)

		return
	}

	h.Metrics.FunctionResponseTime.With(prometheus.Labels{"function": functionName}).Set(float64(time.Since(now).Milliseconds()))

	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}

func (h Handler) Health(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}
