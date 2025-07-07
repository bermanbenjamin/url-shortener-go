package shortener

import "time"

type ShortenRequest struct {
	Code string `json:"code"`
}

type ShortenResponse struct {
	ShortURL string `json:"short_url"`
}

type ShortenURL struct {
	Code      string    `json:"code"`
	URL       string    `json:"url"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}
