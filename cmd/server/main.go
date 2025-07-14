package main

import (
	"github.com/bermanbenjamin/go-shortener-url/internal/shortener"
	"github.com/bermanbenjamin/go-shortener-url/pkg/config"
	"github.com/bermanbenjamin/go-shortener-url/pkg/db"
	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	config := config.NewConfig()
	client := db.InitDatabase(config.Variables.GetString("database.uri"))
	shortenerRepo := shortener.NewShortenerRepository(client)
	shortenerService := shortener.NewShortenerService(shortenerRepo)
	shortenerHandler := shortener.NewShortenerHandler(r, shortenerService)

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.Default())
	gin.SetMode(gin.ReleaseMode)

	r.POST("/shorten", shortenerHandler.Shorten)
	r.GET("/:code", shortenerHandler.Get)

	r.Run(":" + config.Variables.GetString("server.port"))
}
