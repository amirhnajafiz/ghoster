package file

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/google/uuid"
)

const functionsDir = "functions"

func handleUploads(w http.ResponseWriter, r *http.Request) {
	uid := uuid.NewString()
	path := fmt.Sprintf("%s/%s", functionsDir, uid)
	fileName := r.FormValue("file_name")

	file, _, err := r.FormFile("file")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		log.Println(err)

		return
	}

	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
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

	if err := unzip(path, fmt.Sprintf("%s/%s", functionsDir, fileName)); err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		log.Println(err)

		return
	}

	if err := os.RemoveAll(path); err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		log.Println(err)

		return
	}

	w.WriteHeader(http.StatusOK)
}

func health(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}
