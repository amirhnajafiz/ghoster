package filemanager

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/google/uuid"
)

const (
	formFileLabel = "project"
	storageDir    = "/files"
)

// upload a project into files directory.
func (h Handler) upload(w http.ResponseWriter, r *http.Request) {
	// parse our multipart form, 10 << 20 specifies a maximum upload of 10 MB files.
	r.ParseMultipartForm(h.FileLimit << 20)

	// get file from form-data
	file, handler, err := r.FormFile(formFileLabel)
	if err != nil {
		return
	}

	// create the project directory
	path, err := createProjectDir(uuid.NewString())
	if err != nil {
		return
	}

	// create the project meta.json file
	if err := createProjectMetaFile(path, handler.Filename); err != nil {
		return
	}

	// create destination
	dest, err := os.Create(fmt.Sprintf("%s/%s", path, handler.Filename))
	if err != nil {
		return
	}

	// copy the files content
	if _, err := io.Copy(dest, file); err != nil {
		return
	}

	file.Close()
	dest.Close()

	w.Write([]byte("OK"))
}