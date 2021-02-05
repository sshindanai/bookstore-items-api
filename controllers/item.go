package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/sshindanai/repo/bookstore-items-api/services"

	"github.com/sshindanai/bookstore-utils-go/resterrors"

	"github.com/sshindanai/repo/bookstore-items-api/utils/httputils"

	"github.com/sshindanai/bookstore-oauth-go/oauth"
	"github.com/sshindanai/repo/bookstore-items-api/domain/models"
)

var (
	ItemsController itemsControllerInterface = &itemsController{}
)

type itemsControllerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
}

type itemsController struct{}

func (h *itemsController) Create(w http.ResponseWriter, r *http.Request) {
	var item models.Item
	if err := oauth.AuthenticateRequest(r); err != nil {
		// TODO: return err to the caller
		httputils.RespondError(w, err)
		return
	}

	sellerID := oauth.CallerID(r)
	if sellerID == 0 {
		httputils.RespondError(w, resterrors.NewUnauthorizedError("invalid access token"))
		return
	}

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respErr := resterrors.NewBadRequestError(err.Error())
		httputils.RespondError(w, respErr)
		return
	}

	if err := json.Unmarshal(requestBody, &item); err != nil {
		respErr := resterrors.NewBadRequestError("invalid json body")
		httputils.RespondError(w, respErr)
		return
	}

	defer r.Body.Close()

	item.Seller = sellerID
	output := services.ItemsService.Create(&item)
	result := <-output
	if result.Error != nil {
		w.WriteHeader(result.Error.StatusCode())
		return
	}

	httputils.RespondJSON(w, http.StatusCreated, result.Result)
}

func (h *itemsController) Get(w http.ResponseWriter, r *http.Request) {
	if err := oauth.AuthenticateRequest(r); err != nil {
		// TODO: return err to the caller
		httputils.RespondError(w, err)
		return
	}

	httputils.RespondJSON(w, http.StatusOK, r.Header.Get("X-CallerId"))
}
