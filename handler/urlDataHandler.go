package handler

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	mRand "math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/Kleverson-c/jabalinks-backend/service"
	"github.com/Kleverson-c/jabalinks-backend/store"
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

func ShortUrlHandler(writer http.ResponseWriter, request *http.Request) {
	urlData, err := service.GetUrlData(request.PathValue("key"))

	if err != nil {
		http.Error(writer, "failed to get Url data", http.StatusInternalServerError)
		fmt.Print(err.Error())
		return
	}

	if shouldRickRoll() {
		urlData.Url = "https://streamable.com/zxxayc"
	}

	http.Redirect(writer, request, urlData.Url, http.StatusFound)
}

func shouldRickRoll() bool {
	source := mRand.NewSource(time.Now().UnixNano())
	randGemerator := mRand.New(source)
	chance := randGemerator.Intn(1000) + 1

	return chance == 333
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

	err = service.SaveUrlData(uuid, url, createdAt, manualRedirect, redirectText)

	if err != nil {
		http.Error(writer, "failed to save Url data", http.StatusInternalServerError)
		fmt.Print(err.Error())
		return
	}

	json, err := json.Marshal(store.UrlData{Url: stringBuilder.String()})

	if err != nil {
		http.Error(writer, "failed to parse JSON", http.StatusInternalServerError)
		fmt.Print(err.Error())
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	writer.Write(json)
}
