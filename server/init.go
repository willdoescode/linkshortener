package main

import (
	"fmt"
	"log"
	h "main/handlers"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

var (
	CorsOptions = cors.Options{
		AllowedOrigins: []string{"http://localhost:8080"},
	}
)

func serve(port int) {
	router := mux.NewRouter()
	router.HandleFunc("/create", h.CreateShort).Methods("POST")
	router.HandleFunc("/{id}", h.GetShortUrl).Methods("GET")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), cors.New(CorsOptions).Handler(router)))
}
