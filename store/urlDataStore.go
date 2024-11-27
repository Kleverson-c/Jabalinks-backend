package store

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UrlData struct {
	ID             string `json:"id"`
	CreatedAt      string `json:"created_at"`
	Url            string `json:"url"`
	ManualRedirect bool   `json:"manual_redirect"`
	RedirectText   string `json:"redirect_text"`
}

func GetUrlData(id string) (UrlData, error) {
	dbpool, err := pgxpool.New(context.Background(), os.Getenv("DB_URL"))

	if err != nil {
		return UrlData{}, err
	}

	defer dbpool.Close()

	var urlData UrlData
	err = dbpool.QueryRow(context.Background(),
		"SELECT * FROM urls WHERE ID = $1",
		id).Scan(&urlData)

	if err != nil {
		return UrlData{}, err
	}

	return urlData, nil
}

func SaveUrlData(id string, url string, createdAt string, manualRedirect bool, redirectText string) error {
	dbpool, err := pgxpool.New(context.Background(), os.Getenv("DB_URL"))

	if err != nil {
		return err
	}

	defer dbpool.Close()

	_, err = dbpool.Exec(context.Background(),
		"INSERT INTO urls (ID, CreatedAt, Url, ManualRedirect, RedirectText) VALUES ($1, $2, $3, $4, $5)",
		id,
		createdAt,
		url,
		manualRedirect,
		redirectText)

	if err != nil {
		return err
	}

	return nil
}
