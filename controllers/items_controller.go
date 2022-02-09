package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/Bookstore-GolangMS/bookstore_oauth-go/oauth"
	"github.com/LordRadamanthys/bookstore_utils-go/rest_errors"
	"github.com/bookstore_items-api/domain/items"
	"github.com/bookstore_items-api/domain/queries"
	"github.com/bookstore_items-api/services"
	"github.com/bookstore_items-api/utils/http_utils"
	"github.com/gorilla/mux"
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
		restError := rest_errors.NewUnauthorizedError("unauthorized")
		http_utils.RespondError(w, restError)
		return
	}

	//change to unauthorized
	seller := oauth.GetCallerId(r)
	if seller == 0 {
		restError := rest_errors.NewBadRequestError("unauthorized")
		http_utils.RespondError(w, restError)
		return
	}

	var itemRequest items.Item
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		restError := rest_errors.NewBadRequestError("invalid request body")
		http_utils.RespondError(w, restError)
		return
	}
	defer r.Body.Close()

	if err := json.Unmarshal(requestBody, &itemRequest); err != nil {
		restError := rest_errors.NewBadRequestError("invalid json body")
		http_utils.RespondError(w, restError)
		return
	}

	itemRequest.Seller = int(seller)
	result, createErr := services.ItemsService.Create(itemRequest)

	if createErr != nil {
		http_utils.RespondError(w, createErr)
		return
	}

	http_utils.RespondJson(w, http.StatusCreated, result)
}

func (c *itemsControllerStruct) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	itemId := strings.TrimSpace(vars["id"])

	item, err := services.ItemsService.Get(itemId)
	if err != nil {
		http_utils.RespondError(w, err)
		return
	}
	http_utils.RespondJson(w, http.StatusOK, item)
}

func (c *itemsControllerStruct) Search(w http.ResponseWriter, r *http.Request) {

	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		apiErr := rest_errors.NewBadRequestError("ivalid json body")
		http_utils.RespondError(w, apiErr)
		return
	}
	defer r.Body.Close()

	var query queries.EsQuery
	if err := json.Unmarshal(bytes, &query); err != nil {
		apiErr := rest_errors.NewBadRequestError("invalid json body")
		http_utils.RespondError(w, apiErr)
		return
	}

	items, searchErr := services.ItemsService.Search(query)
	if err != nil {
		http_utils.RespondError(w, searchErr)
		return
	}
	http_utils.RespondJson(w, http.StatusOK, items)
}
