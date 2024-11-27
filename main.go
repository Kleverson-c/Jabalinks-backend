package main

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	mRand "math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var urlMap = make(map[string]urlData)

type urlData struct {
	URL               string    `json:"url"`
	Date              time.Time `json:"date"`
	NeedsConfirmation bool      `json:"needsConfirmation"`
}

func enableCors(writer *http.ResponseWriter) {
	(*writer).Header().Set("Access-Control-Allow-Origin", "*")
}

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
	urlData := urlMap[request.PathValue("key")]
	fmt.Print(urlData)
	if urlData.URL == "" {
		http.Error(writer, "Page not found", http.StatusNotFound)
	}

	if shouldRickRoll() {
		urlData.URL = "https://streamable.com/zxxayc"
	}

	http.Redirect(writer, request, urlData.URL, http.StatusFound)
}

func urlHandler(writer http.ResponseWriter, request *http.Request) {
	enableCors(&writer)
	url := request.FormValue("url")
	needsConfirmation, _ := strconv.ParseBool(request.FormValue("needsConfirmation"))
	hostName := request.Host

	var stringBuilder strings.Builder
	stringBuilder.WriteString("http://")
	stringBuilder.WriteString(hostName)
	stringBuilder.WriteString("/short/")

	uuid := generateUUID()
	stringBuilder.WriteString(uuid)

	urlMap[uuid] = urlData{URL: url, NeedsConfirmation: needsConfirmation, Date: time.Now()}
	json, error := json.Marshal(urlData{URL: stringBuilder.String(), NeedsConfirmation: needsConfirmation})

	if error != nil {
		http.Error(writer, "Failed to parse JSON", http.StatusInternalServerError)
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.Write(json)
}

func shouldRickRoll() bool {
	source := mRand.NewSource(time.Now().UnixNano())
	randGemerator := mRand.New(source)
	chance := randGemerator.Intn(1000) + 1

	return chance == 333
}
