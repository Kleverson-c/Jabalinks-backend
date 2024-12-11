package handler

import (
	mRand "math/rand"
	"net/http"
	"time"
)

func ShortUrlHandler(writer http.ResponseWriter, request *http.Request) {
	urlData := UrlMap[request.PathValue("key")]

	if urlData.Url == "" {
		http.Error(writer, "failed to get Url data", http.StatusInternalServerError)
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
