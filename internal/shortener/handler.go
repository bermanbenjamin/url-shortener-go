package shortener

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ShortenerHandler struct {
	gin     *gin.Engine
	service *ShortenerService
}

func NewShortenerHandler(gin *gin.Engine, service *ShortenerService) *ShortenerHandler {
	return &ShortenerHandler{gin: gin, service: service}
}

func (h *ShortenerHandler) Shorten(c *gin.Context) {
	var shortenUrl ShortenURL
	if err := c.ShouldBindJSON(&shortenUrl); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

}
