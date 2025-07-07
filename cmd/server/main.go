package main

import (
	"github.com/bermanbenjamin/go-shortener-url/config"
	"github.com/bermanbenjamin/go-shortener-url/internal/shortener"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	config := config.NewConfig()
	shortenerRepo := shortener.NewShortenerRepository(config.Client)
	shortenerService := shortener.NewShortenerService(shortenerRepo)
	shortenerHandler := shortener.NewShortenerHandler(r, shortenerService)

	r.POST("/shorten", shortenerHandler.Shorten)
	r.GET("/:code", shortenerHandler.Get)

	r.Run(":" + config.Variables.GetString("server.port"))
}
