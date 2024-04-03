package file

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/google/uuid"
)

func handleUploads(functionsDir, prefixToken string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		uid := uuid.NewString()
		path := fmt.Sprintf("%s/%s", functionsDir, uid)
		fileName := r.FormValue("file_name")
		newPath := fmt.Sprintf("%s/%s%s", functionsDir, prefixToken, fileName)

		file, _, err := r.FormFile("file")
		if err != nil {
			errorHandler(w, http.StatusBadRequest, err)
			return
		}

		f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0777)
		if err != nil {
			errorHandler(w, http.StatusInternalServerError, err)
			return
		}

		if _, err := io.Copy(f, file); err != nil {
			errorHandler(w, http.StatusInternalServerError, err)
			return
		}

		_ = file.Close()
		_ = f.Close()

		if err := os.Mkdir(newPath, 0777); err != nil && !errors.Is(err, os.ErrExist) {
			errorHandler(w, http.StatusInternalServerError, err)
			return
		}

		if err := unzip(path, newPath); err != nil {
			errorHandler(w, http.StatusInternalServerError, err)
			return
		}

		if err := os.RemoveAll(path); err != nil {
			errorHandler(w, http.StatusInternalServerError, err)
			return
		}

		log.Printf("file-server [%s] %s\n", r.Method, r.RequestURI)

		w.WriteHeader(http.StatusOK)
	}
}

func health(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func errorHandler(w http.ResponseWriter, status int, err error) {
	log.Printf("ghoster handler error: %v\n", err)
	w.WriteHeader(status)
}
