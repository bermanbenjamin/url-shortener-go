package main

import (
	"github.com/bermanbenjamin/go-shortener-url/internal/shortener"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	shortenerHandler := shortener.NewShortenerHandler(r)

	r.POST("/shorten", shortenerHandler.Shorten)

	r.Run(":8080")
}
