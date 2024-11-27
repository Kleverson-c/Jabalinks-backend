package main

import (
	"log"
	"net/http"

	"github.com/Kleverson-c/jabalinks-backend/handler"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /short/{key}", handler.ShortUrlHandler)
	mux.HandleFunc("POST /url", handler.UrlHandler)

	log.Fatal(http.ListenAndServe(":8080", mux))
}
