package controllers

import (
	"fmt"
	"net/http"

	"github.com/Bookstore-GolangMS/bookstore_oauth-go/oauth"
	"github.com/bookstore_items-api/domain/items"
	"github.com/bookstore_items-api/services"
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
		return
	}

	item := items.Item{
		Seller: int(oauth.GetCallerId(r)),
	}

	result, err := services.ItemsService.Create(item)
	if err.Message != "" {
		return
	}

	fmt.Println(result)
}

func (c *itemsControllerStruct) Get(w http.ResponseWriter, r *http.Request) {

}
