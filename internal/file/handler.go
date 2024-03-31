package file

import "net/http"

func handleUploads(w http.ResponseWriter, r *http.Request) {

}

func health(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}
