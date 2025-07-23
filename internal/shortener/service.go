package shortener

import (
	"context"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type ShortenerService interface {
	Shorten(shortenUrl ShortenURL) (string, error)
	Get(code string) (string, error)
}

type shortenerService struct {
	repo        ShortenerRepository
	redisClient *redis.Client
}

func NewShortenerService(repo ShortenerRepository, redisClient *redis.Client) ShortenerService {
	return &shortenerService{repo: repo, redisClient: redisClient}
}

func (s *shortenerService) Shorten(shortenUrl ShortenURL) (string, error) {
	shortenUrl.Code = uuid.New().String()
	return s.repo.Create(shortenUrl)
}

func (s *shortenerService) Get(code string) (string, error) {
	s.redisClient.Get(context.Background(), "")
	return s.repo.Get(code)
}
