package shortener

import (
	"context"
	"log"
	"time"

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
	val, err := s.redisClient.Get(context.Background(), code).Result()
	log.Println("Getting from redis: ", val)
	if err == redis.Nil {
		log.Println("Getting from repo: ", val)
		val, err = s.repo.Get(code)
		if err != nil {
			return "", err
		}
		s.redisClient.Set(context.Background(), code, val, 2*time.Hour)
	} else if err != nil {
		return "", err
	}

	return val, nil
}
