package main

import (
	"log"
	"net/http"
	"url-shortener/handlers"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/shorten", handlers.ShortenURL).Methods("POST")
	router.HandleFunc("/{shortURL}", handlers.RedirectURL).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}
