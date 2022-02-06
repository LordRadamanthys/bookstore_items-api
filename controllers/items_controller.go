package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Bookstore-GolangMS/bookstore_oauth-go/oauth"
	"github.com/LordRadamanthys/bookstore_utils-go/rest_errors"
	"github.com/bookstore_items-api/domain/items"
	"github.com/bookstore_items-api/services"
	"github.com/bookstore_items-api/utils/http_utils"
)

var (
	ItemsController itemsControllerInterface = &itemsControllerStruct{}
)

type itemsControllerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
}

type itemsControllerStruct struct{}

func (c *itemsControllerStruct) Create(w http.ResponseWriter, r *http.Request) {
	if err := oauth.AuthenticateRequest(r); err != nil {
		http_utils.RespondError(w, rest_errors.RestErr{})
		return
	}

	var itemRequest items.Item
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		restError := rest_errors.BadRequestError("invalid request body", err)
		http_utils.RespondError(w, *restError)
		return
	}
	defer r.Body.Close()

	if err := json.Unmarshal(requestBody, &itemRequest); err != nil {
		restError := rest_errors.BadRequestError("invalid json body", err)
		http_utils.RespondError(w, *restError)
		return
	}

	itemRequest.Seller = int(oauth.GetClientId(r))
	result, createErr := services.ItemsService.Create(itemRequest)

	if createErr != nil {
		http_utils.RespondError(w, *createErr)
		return
	}

	http_utils.RespondJson(w, http.StatusCreated, result)
}

func (c *itemsControllerStruct) Get(w http.ResponseWriter, r *http.Request) {

}
