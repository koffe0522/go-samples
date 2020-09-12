package routes

import (
	"koffe0522/go-projects/go-rest/controllers"

	"github.com/gorilla/mux"
)

// NewRouter is Routes setting
func NewRouter() *mux.Router {
	bookController := controllers.NewBookInstance()
	r := mux.NewRouter()
	r.HandleFunc("/books", bookController.Index).Methods("GET")
	r.HandleFunc("/books/{id}", bookController.Show).Methods("GET")
	r.HandleFunc("/books", bookController.Store).Methods("POST")

	return r
}
