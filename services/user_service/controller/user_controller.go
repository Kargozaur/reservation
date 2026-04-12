package controller

import (
	"log/slog"
	"user-service/auth/pass"
	"user-service/auth/token"
	"user-service/handlers"
	"user-service/middleware"
	"user-service/repositories/refresh"
	"user-service/repositories/user"
	"user-service/services/users"
	"user-service/validators/credential"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func buildDeps(db *gorm.DB) *handlers.UserHandler {
	userRepo := user.NewUserRepository(db)
	tokenRepo := refresh.NewRefreshRepository(db)
	phasher := pass.NewHasher(12)
	jwt := token.NewJWT()
	validator := credential.NewValidator()
	service := users.NewUserService(phasher, jwt, userRepo, tokenRepo, validator)
	handler := handlers.NewUserHandler(*service)
	return handler
}

func UserRouter(rg *gin.RouterGroup, db *gorm.DB, logger *slog.Logger) {
	handler := buildDeps(db)
	user := rg.Group("/user")
	{
		user.POST("/register", handler.CreateUser())
		user.POST("/login", handler.LoginUser())

		auth := user.Group("/")
		auth.Use(middleware.VerifyAccess(logger))
		{
			auth.GET("/profile", handler.GetUser())
			auth.PUT("/profile/update_name", handler.UpdateName())
			auth.PUT("/profile/update_email", handler.UpdateEmail())
			auth.PUT("/profile/update_password", handler.UpdatePassword())
			auth.POST("/logout", handler.LogoutHandler())
		}
	}
}
