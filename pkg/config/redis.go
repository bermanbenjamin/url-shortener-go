package config

import (
	"github.com/redis/go-redis/v9"
)

type RedisConfig struct {
	RedisClient *redis.Client
}

func SetupRedis(pass string, addr string) (client *redis.Client) {
	client = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pass, // No password set
		DB:       0,    // Use default DB
		Protocol: 2,    // Connection protocol
	})

	return client
}
