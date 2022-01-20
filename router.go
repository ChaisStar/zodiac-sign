package main

import (
	"asia/handlers"

	"github.com/gorilla/mux"
)

func router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.Handler).
		Methods("POST").
		Schemes("http")
	return r
}
