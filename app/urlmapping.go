package app

import (
	"net/http"

	"github.com/sshindanai/repo/bookstore-items-api/controllers"
)

func mapUrls() {
	r.HandleFunc("/items", controllers.ItemsController.Create).Methods(http.MethodPost)
	r.HandleFunc("/items/{id}", controllers.ItemsController.Get).Methods(http.MethodGet)
}
