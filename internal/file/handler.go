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

const (
	functionsDir = "functions"
	prefixToken  = "xxx-"
)

func handleUploads(w http.ResponseWriter, r *http.Request) {
	uid := uuid.NewString()
	path := fmt.Sprintf("%s/%s", functionsDir, uid)
	fileName := r.FormValue("file_name")
	newPath := fmt.Sprintf("%s/%s%s", functionsDir, prefixToken, fileName)

	file, _, err := r.FormFile("file")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		log.Println(err)

		return
	}

	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		log.Println(err)

		return
	}

	if _, err := io.Copy(f, file); err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		log.Println(err)

		return
	}

	file.Close()
	f.Close()

	if err := os.Mkdir(newPath, 0777); err != nil && !errors.Is(err, os.ErrExist) {
		w.WriteHeader(http.StatusInternalServerError)

		log.Println(err)

		return
	}

	if err := unzip(path, newPath); err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		log.Println(err)

		return
	}

	if err := os.RemoveAll(path); err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		log.Println(err)

		return
	}

	log.Printf("file-server [%s] %s\n", r.Method, r.RequestURI)

	w.WriteHeader(http.StatusOK)
}

func health(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}
