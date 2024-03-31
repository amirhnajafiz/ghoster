package file

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func NewServer(port int) {
	go func() {
		router := mux.NewRouter()

		router.HandleFunc("/files", handleUploads).Methods(http.MethodPost)

		srv := &http.Server{
			Handler: router,
			Addr:    fmt.Sprintf("127.0.0.1:%d", port),
		}

		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()
}
