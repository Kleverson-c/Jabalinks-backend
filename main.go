package main

import (
	"log"
	"net/http"

	"Jabalinks/handler"
)

func main() {
	handler.CreateUrlMap()

	mux := http.NewServeMux()
	mux.HandleFunc("GET /short/{key}", handler.ShortUrlHandler)
	mux.HandleFunc("POST /url", handler.UrlHandler)

	log.Fatal(http.ListenAndServe(":8080", mux))
}
