package main

import (
	"github.com/ChaisStar/zodiac-sign/handlers"

	"github.com/gorilla/mux"
)

func router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api", handlers.Handler).
		Methods("POST").
		Schemes("http")
	spa := spaHandler{staticPath: "frontend", indexPath: "index.html"}
	r.PathPrefix("/").Handler(spa)
	return r
}
