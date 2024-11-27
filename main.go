package main

import (
	"log"
	"net/http"

	"github.com/Kleverson-c/jabalinks-backend/handler"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	mux := http.NewServeMux()
	mux.HandleFunc("GET /short/{key}", handler.ShortUrlHandler)
	mux.HandleFunc("POST /url", handler.UrlHandler)

	log.Fatal(http.ListenAndServe(":8080", mux))
}
