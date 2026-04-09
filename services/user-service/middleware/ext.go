package middleware

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
)

func extractToken(ctx *gin.Context) (string, error) {
	if token, err := ctx.Cookie("access_token"); err == nil {
		return token, nil
	}
	auth := ctx.GetHeader("Authorization")
	if token, ok := strings.CutPrefix(auth, "Bearer "); ok && token != "" {
		return token, nil
	}
	return "", errors.New("Token is not provided")
}
