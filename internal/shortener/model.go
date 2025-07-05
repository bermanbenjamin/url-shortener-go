package shortener

import "time"

type ShortenRequest struct {
	URL string `json:"url"`
}

type ShortenResponse struct {
	ShortURL string `json:"short_url"`
}

type ShortenURL struct {
	ID        string    `json:"id"`
	Code      string    `json:"code"`
	URL       string    `json:"url"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}
