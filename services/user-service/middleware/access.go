package middleware

import (
	"log/slog"
	"net/http"
	"user-service/auth/token"

	"github.com/gin-gonic/gin"
)

func VerifyAccess(logger *slog.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		jwt := token.NewJWT()
		if token, err := extractToken(ctx); err == nil {
			if userID, err := jwt.GetUId(token); err == nil {
				ctx.Set("userID", userID)
				logger.Info("Parsed JWT with middleware")
				ctx.Next()
				return
			}
		}
		logger.Info("Failed to parse JWT with middleware")
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}
}
