package filemanager

import (
	"io/ioutil"
	"net/http"
)

func (h Handler) upload(w http.ResponseWriter, r *http.Request) {
	// parse our multipart form, 10 << 20 specifies a maximum upload of 10 MB files.
	r.ParseMultipartForm(h.FileLimit << 20)

	file, handler, err := r.FormFile("project")
	if err != nil {
		return
	}

	defer file.Close()

	tempFile, err := ioutil.TempFile("temp-images", "upload-*.png")
	if err != nil {
		return
	}

	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return
	}

	// write this byte array to our temporary file
	tempFile.Write(fileBytes)
}
