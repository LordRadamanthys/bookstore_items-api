package app

import (
	"net/http"

	"github.com/bookstore_items-api/controllers"
	"github.com/gorilla/mux"
)

var (
	router = mux.NewRouter()
)

func MapUrls() {
	router.HandleFunc("/items", controllers.ItemsController.Create).Methods(http.MethodPost)
	router.HandleFunc("/ping", controllers.PinController.Ping).Methods(http.MethodGet)
	router.HandleFunc("/items/{id}", controllers.ItemsController.Get).Methods(http.MethodGet)
}
