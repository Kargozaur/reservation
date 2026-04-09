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
		token, err := extractToken(ctx)
		if err != nil {
			logger.Info("Failed to parse JWT with middleware(Did not extract token)")
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		userID, err := jwt.GetUId(token)
		if err != nil {
			logger.Info("Failed to parse JWT with middleware(Failed to get id)")
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		ctx.Set("userID", userID)
		logger.Info("Parsed JWT with middleware")
		ctx.Next()
	}
}
