package shortener

import "time"

type ShortenResponse struct {
	Url string `json:"url"`
}

type ShortenURL struct {
	Code      string    `json:"code"`
	URL       string    `json:"url"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}
