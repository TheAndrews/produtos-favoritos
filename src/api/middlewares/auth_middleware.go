package middlewares

import (
	"net/http"
	"produtos-favoritos/src/infrastructure/config"

	"github.com/gin-gonic/gin"
)

func APIKeyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.GetHeader("x-api-key")

		if apiKey == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing API key"})
			return
		}

		if apiKey != config.API_KEY {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid API key"})
			return
		}

		c.Set("authenticated", true)
		c.Next()
	}
}
