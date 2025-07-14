package config

import (
	"github.com/redis/go-redis/v9"
)

type RedisConfig struct {
	RedisClient *redis.Client
}

func (r *RedisConfig) SetupRedis() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // No password set
		DB:       0,  // Use default DB
		Protocol: 2,  // Connection protocol
	})

	r.RedisClient = client

}
