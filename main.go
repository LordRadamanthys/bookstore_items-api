package main

import (
	"github.com/bookstore_items-api/app"
	"github.com/gorilla/mux"
)

var (
	router = mux.NewRouter()
)

func main() {
	app.StartApplication()
}
