package shortener

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ShortenerHandler struct {
	gin *gin.Engine
}

func NewShortenerHandler(gin *gin.Engine) *ShortenerHandler {
	return &ShortenerHandler{gin: gin}
}

func (h *ShortenerHandler) Shorten(c *gin.Context) {
	var shortenUrl ShortenURL
	if err := c.ShouldBindJSON(&shortenUrl); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

}
