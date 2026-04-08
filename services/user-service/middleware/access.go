package middleware

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

func VerifyAccess(logger *slog.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, err := ctx.Cookie("access_token")
		if err != nil {
			logger.Error("Failed to get access token cookie: " + err.Error())
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		logger.Info("User have access token")
		ctx.Next()
	}
}
