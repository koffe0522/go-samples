package controllers

import (
	"encoding/json"
	"koffe0522/go-projects/go-rest/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// BookController struct
type BookController struct{}

// NewBookInstance constructer
func NewBookInstance() *BookController {
	return new(BookController)
}

// Index route
func (c *BookController) Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		models.All(),
	))
}

// Show route
func (c *BookController) Show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	strID := vars["id"]
	id, _ := strconv.Atoi(strID)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		models.FindByID(id),
	))
}

// Store route
func (c *BookController) Store(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		json.NewEncoder(w).Encode("bad request")
	}

	book := models.Book{}
	book.Title = r.FormValue("title")
	result := models.Create(book)

	w.Header().Set("Content-Type", "application/json")
	if result {
		json.NewEncoder(w).Encode(newResponse(
			http.StatusCreated,
			http.StatusText(http.StatusCreated),
			"Success",
		))
		return
	}

	json.NewEncoder(w).Encode(newResponse(
		http.StatusInternalServerError,
		http.StatusText(http.StatusInternalServerError),
		"Faild",
	))
}
