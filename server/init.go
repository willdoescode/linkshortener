package main

import (
	"log"
	h "main/handlers"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

var (
	CorsOptions = cors.Options{
		AllowedOrigins: []string{"http://localhost:8081"},
	}
)

func serve() {
	router := mux.NewRouter()
	router.HandleFunc("/create", h.CreateShort).Methods("POST")
	router.HandleFunc("/{id}", h.GetShortUrl).Methods("GET")
	handler := cors.New(CorsOptions).Handler(router)

	log.Fatal(http.ListenAndServe(":8080", handler))
}
