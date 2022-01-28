package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/handlers"
)

func main() {
	log.Print("http server started on [::]:9100")

	headersOk := handlers.AllowedHeaders([]string{"Origin, X-Requested-With, Content-Type, Accept, x-client-key, x-client-token, x-client-secret, Authorization", "Content-Disposition"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	originsOk := handlers.AllowedOrigins([]string{"*"})

	router := router()

	srv := &http.Server{
		Handler:      handlers.CORS(originsOk, headersOk, methodsOk)(router),
		Addr:         ":9100",
		WriteTimeout: 120 * time.Second,
		ReadTimeout:  120 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
