package app

import (
	"net/http"
	"time"
)

func StartApplication() {
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
