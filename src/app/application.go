package app

import (
	"log"
	"net/http"
)

func StartApplication() {
	mapUrls()

	log.Print("start the application...")

	srv := http.Server{
		Addr:    ":8020",
		Handler: router,
	}
	srv.ListenAndServe()
}
