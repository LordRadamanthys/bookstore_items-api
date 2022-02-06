package services

import (
	"net/http"

	"github.com/LordRadamanthys/bookstore_utils-go/rest_errors"
	"github.com/bookstore_items-api/domain/items"
)

var (
	ItemsService itemsServiceInterface = &itemsService{}
)

type itemsServiceInterface interface {
	Create(items.Item) (*items.Item, *rest_errors.RestErr)
	Get(string) (*items.Item, rest_errors.RestErr)
}

type itemsService struct{}

func (s *itemsService) Create(item items.Item) (*items.Item, *rest_errors.RestErr) {
	return nil, &rest_errors.RestErr{
		Message: "Not implemmented",
		Status:  http.StatusNotImplemented,
		Error:   "",
		Causes:  nil,
	}
}

func (s *itemsService) Get(id string) (*items.Item, rest_errors.RestErr) {
	return nil, rest_errors.RestErr{
		Message: "Not implemmented",
		Status:  http.StatusNotImplemented,
		Error:   "",
		Causes:  nil,
	}
}
