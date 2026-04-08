package middleware

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

func VerifyAccess(logger *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := c.Cookie("access_token")
		if err != nil {
			logger.Error("Failed to get access token cookie: " + err.Error())
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Next()
		logger.Info("User have access token")
	}
}
