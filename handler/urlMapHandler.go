package handler

var UrlMap map[string]UrlData

type UrlData struct {
	ID             string `json:"id"`
	CreatedAt      string `json:"created_at"`
	Url            string `json:"url"`
	ManualRedirect bool   `json:"manual_redirect"`
	RedirectText   string `json:"redirect_text"`
}

func CreateUrlMap() {
	UrlMap = make(map[string]UrlData)
}
