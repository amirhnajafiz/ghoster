package router

import (
	"github.com/gorilla/mux"

	"restapi/restapi/internal/routes/destroy"
	"restapi/restapi/internal/routes/index"
	"restapi/restapi/internal/routes/show"
	"restapi/restapi/internal/routes/store"
	"restapi/restapi/internal/routes/update"
)

// GetRouter : returns the application router
func GetRouter() *mux.Router {
	// Init router
	r := mux.NewRouter()

	// Route handlers / Endpoints
	r.HandleFunc("/api/books", index.GetBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", show.GetBook).Methods("GET")
	r.HandleFunc("/api/books", store.CreateBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", update.UpdateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", destroy.DeleteBook).Methods("DELETE")

	return r
}
