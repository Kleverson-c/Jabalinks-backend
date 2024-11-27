package service

import "github.com/Kleverson-c/jabalinks-backend/store"

func GetUrlData(id string) (store.UrlData, error) {
	return store.GetUrlData(id)
}

func SaveUrlData(id string, url string, createdAt string, manualRedirect bool, redirectText string) error {
	return store.SaveUrlData(id, url, createdAt, manualRedirect, redirectText)
}
