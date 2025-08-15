package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthTokenMiddleware(authToken string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Header.Get("X-Auth-Token") != authToken {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Invalid auth token"})
			return
		}

		c.Next()
	}
}
