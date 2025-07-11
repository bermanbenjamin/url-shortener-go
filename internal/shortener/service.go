package shortener

import "github.com/google/uuid"

type ShortenerService interface {
	Shorten(shortenUrl ShortenURL) (string, error)
	Get(code string) (string, error)
}

type shortenerService struct {
	repo ShortenerRepository
}

func NewShortenerService(repo ShortenerRepository) ShortenerService {
	return &shortenerService{repo: repo}
}

func (s *shortenerService) Shorten(shortenUrl ShortenURL) (string, error) {
	shortenUrl.Code = uuid.New().String()[:6]
	return s.repo.Create(shortenUrl)
}

func (s *shortenerService) Get(code string) (string, error) {
	return s.repo.Get(code)
}
