package main

import (
	"github.com/bermanbenjamin/go-shortener-url/internal/shortener"
	"github.com/bermanbenjamin/go-shortener-url/pkg/config"
	"github.com/bermanbenjamin/go-shortener-url/pkg/db"
	"github.com/gin-contrib/cors"
	"github.com/spf13/viper"

	"github.com/gin-gonic/gin"
)

func main() {
	viper.SetConfigFile(".env")
	_ = viper.ReadInConfig() // Ignore error if .env does not exist
	viper.AutomaticEnv()     // Ensure viper loads environment variables

	r := gin.Default()
	mongoURI := viper.GetString("MONGO_URI")
	client := db.InitDatabase(mongoURI)

	redisAddr := viper.GetString("REDIS_ADDR")
	redisPass := viper.GetString("REDIS_PASS")
	redisClient := config.SetupRedis(redisPass, redisAddr)

	shortenerRepo := shortener.NewShortenerRepository(client)
	shortenerService := shortener.NewShortenerService(shortenerRepo, redisClient)
	shortenerHandler := shortener.NewShortenerHandler(r, shortenerService)

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.Default())
	gin.SetMode(gin.ReleaseMode)

	r.POST("/shorten", shortenerHandler.Shorten)
	r.GET("/:code", shortenerHandler.Get)

	r.Run(":" + viper.GetString("SERVER_PORT"))
}
