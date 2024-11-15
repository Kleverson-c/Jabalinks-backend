package main

import (
	"crypto/rand"
	"encoding/base64"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var urlMap = make(map[string]string)

func generateUUID() string {
	uuid := make([]byte, 16)
	_, error := rand.Read(uuid)
	if error != nil {
		panic(error)
	}

	return string(base64.RawURLEncoding.EncodeToString(uuid))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /short/{key}", shortUrlHandler)
	mux.HandleFunc("POST /url", urlHandler)

	log.Fatal(http.ListenAndServe(":8080", mux))
}

func shortUrlHandler(writer http.ResponseWriter, request *http.Request) {
	url := urlMap[request.PathValue("key")]
	if url == "" {
		http.Error(writer, "Page not found", http.StatusNotFound)
	}

	http.Redirect(writer, request, url, http.StatusFound)
}

func urlHandler(writer http.ResponseWriter, request *http.Request) {
	url := request.FormValue("url")
	isNsfw, parseError := strconv.ParseBool(request.FormValue("nsfw"))
	hostName := request.Host

	var stringBuilder strings.Builder
	stringBuilder.WriteString(hostName)
	stringBuilder.WriteString("/short/")
	if parseError == nil && isNsfw {

	}

	uuid := generateUUID()
	stringBuilder.WriteString(uuid)
	urlMap[uuid] = url

	writer.Write([]byte(stringBuilder.String()))
}
