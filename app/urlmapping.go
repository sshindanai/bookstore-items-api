package app

import (
	"net/http"

	"github.com/sshindanai/repo/bookstore-items-api/controllers"
)

var (
	handler = controllers.NewItemHandler()
)

func mapUrls() {
	r.HandleFunc("/items", handler.Create).Methods(http.MethodPost)
	r.HandleFunc("/items/{id}", handler.Get).Methods(http.MethodGet)
}
