package handlers

import (
	"errors"
	"net/http"
	"user-service/schemas/request"
	"user-service/services/users"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type UserHandler struct {
	service users.UserService
}

func NewUserHandler(service users.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (c *UserHandler) CreateUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var newUser request.RegisterSchema
		if err := ctx.ShouldBindJSON(&newUser); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		user, err := c.service.RegisterUser(newUser)
		if err != nil {
			ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusCreated, user)
	}
}

func (c *UserHandler) LoginUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var loginUser request.LoginSchema
		if err := ctx.ShouldBindJSON(&loginUser); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		accessToken, refreshToken, err := c.service.LoginUser(loginUser)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.SetCookie("access_token", accessToken, 1800, "/", "localhost", false, true)
		ctx.SetCookie("refresh_token", refreshToken, 60*60*24*7, "/", "localhost", false, true)
		ctx.JSON(http.StatusCreated, gin.H{"message": "login successful"})
	}
}

func (c *UserHandler) GetUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cookie, err := ctx.Cookie("access_token")
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
		user, err := c.service.GetUser(cookie)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, user)
	}
}

func (c *UserHandler) UpdateName() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var updateName request.UpdateNameSchema
		if err := ctx.BindJSON(&updateName); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if updateName.FirstName == nil && updateName.LastName == nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "at least one field must be provided"})
			return
		}
		cookie, err := ctx.Cookie("access_token")
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
		if err := c.service.UpdateName(cookie, updateName); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "name updated"})
	}
}

func (c *UserHandler) UpdateEmail() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var updateEmail request.UpdateEmailSchema
		if err := ctx.ShouldBindJSON(&updateEmail); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		cookie, err := ctx.Cookie("access_token")
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
		if err := c.service.UpdateEmail(cookie, updateEmail); err != nil {
			if errors.Is(err, jwt.ErrTokenExpired) {
				ctx.JSON(http.StatusUnauthorized, gin.H{"error": "token expired"})
				return
			}
			if errors.Is(err, gorm.ErrDuplicatedKey) {
				ctx.JSON(http.StatusConflict, gin.H{"error": "email already exists"})
				return
			}
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "email updated"})
	}
}

func (c *UserHandler) UpdatePassword() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var updatePassword request.UpdatePasswordSchema
		if err := ctx.ShouldBindJSON(&updatePassword); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		cookie, err := ctx.Cookie("access_token")
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
		if err := c.service.UpdatePassword(cookie, updatePassword); err != nil {
			if errors.Is(err, jwt.ErrTokenExpired) {
				ctx.JSON(http.StatusUnauthorized, gin.H{"error": "token expired"})
				return
			}
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "password updated"})
	}
}
