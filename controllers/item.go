package controllers

import (
	"net/http"

	"github.com/sshindanai/repo/bookstore-items-api/utils/httputils"

	"github.com/sshindanai/bookstore-oauth-go/oauth"
	"github.com/sshindanai/repo/bookstore-items-api/domain/models"
	"github.com/sshindanai/repo/bookstore-items-api/services"
)

type ItemsHandlerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
}

type itemsHandler struct {
	Service services.ItemsServiceInterface
}

func NewItemHandler() ItemsHandlerInterface {
	return &itemsHandler{
		Service: services.NewItemService(),
	}
}

func (h *itemsHandler) Create(w http.ResponseWriter, r *http.Request) {
	if err := oauth.AuthenticateRequest(r); err != nil {
		// TODO: return err to the caller
		httputils.RespondJSON(w, err.Code, err)
		return
	}
	item := models.Item{
		Seller: oauth.CallerID(r),
	}
	output := make(chan models.ItemConcurrent)

	go h.Service.Create(&item, output)
	result := <-output

	if result.Error != nil {
		w.WriteHeader(result.Error.Code)
		return
	}
	httputils.RespondJSON(w, http.StatusCreated, result.Result)
}

func (h *itemsHandler) Get(w http.ResponseWriter, r *http.Request) {
	if err := oauth.AuthenticateRequest(r); err != nil {
		// TODO: return err to the caller
		httputils.RespondJSON(w, err.Code, err)
		return
	}

	httputils.RespondJSON(w, http.StatusOK, "Welcome")
}
