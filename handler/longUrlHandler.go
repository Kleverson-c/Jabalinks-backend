package handler

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func EnableCors(writer *http.ResponseWriter) {
	(*writer).Header().Set("Access-Control-Allow-Origin", "*")
}

func generateUUID() (string, error) {
	uuid := make([]byte, 16)
	_, err := rand.Read(uuid)
	if err != nil {
		return "", err
	}

	return string(base64.RawURLEncoding.EncodeToString(uuid)), nil
}

func UrlHandler(writer http.ResponseWriter, request *http.Request) {
	EnableCors(&writer)
	url := request.FormValue("url")

	if url == "" {
		http.Error(writer, "Missing URL", http.StatusBadRequest)
		return
	}

	manualRedirect, _ := strconv.ParseBool(request.FormValue("manual_redirect"))
	createdAt := time.Now().String()
	redirectText := request.FormValue("redirect_text")

	hostName := request.Host
	var stringBuilder strings.Builder
	stringBuilder.WriteString("http://")
	stringBuilder.WriteString(hostName)
	stringBuilder.WriteString("/short/")

	uuid, err := generateUUID()

	if err != nil {
		http.Error(writer, "failed to generate ID", http.StatusInternalServerError)
		fmt.Print(err.Error())
		return
	}

	stringBuilder.WriteString(uuid)

	UrlMap[uuid] = UrlData{
		ID:             uuid,
		CreatedAt:      createdAt,
		Url:            url,
		ManualRedirect: manualRedirect,
		RedirectText:   redirectText,
	}

	json, err := json.Marshal(UrlData{ID: stringBuilder.String()})

	if err != nil {
		http.Error(writer, "failed to parse JSON", http.StatusInternalServerError)
		fmt.Print(err.Error())
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	writer.Write(json)
}
