package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	log.Print("http server started on [::]:9100")
	router := router()
	srv := &http.Server{
		Handler:      router,
		Addr:         ":9100",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
