package app

import (
	"net/http"
	"time"

	"github.com/bookstore_items-api/clients/elasticsearch"
)

func StartApplication() {
	elasticsearch.Init()
	MapUrls()

	server := &http.Server{
		Handler:      router,
		Addr:         "localhost:8083",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
