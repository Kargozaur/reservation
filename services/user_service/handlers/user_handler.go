package handlers

import (
	"errors"
	"net/http"
	"user-service/schemas/request"
	"user-service/services/users"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
			if errors.Is(err, gorm.ErrDuplicatedKey) {
				ctx.JSON(http.StatusConflict, gin.H{"error": "user already exists"})
				return
			}
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
		accessToken, refreshToken, err := c.service.LoginUser(&loginUser)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				ctx.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
				return
			}
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
		userID, ok := ctx.Get("userID")
		if !ok {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
		user, err := c.service.GetUser(userID.(uuid.UUID))
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
		userID, ok := ctx.Get("userID")
		if !ok {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
		if err := c.service.UpdateName(userID.(uuid.UUID), updateName); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
		userID, ok := ctx.Get("userID")
		if !ok {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
		if err := c.service.UpdateEmail(userID.(uuid.UUID), updateEmail); err != nil {
			if errors.Is(err, gorm.ErrDuplicatedKey) {
				ctx.JSON(http.StatusConflict, gin.H{"error": "email is taken"})
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
		userID, ok := ctx.Get("userID")
		if !ok {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
		if err := c.service.UpdatePassword(userID.(uuid.UUID), updatePassword); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "password updated"})
	}
}

func (c *UserHandler) LogoutHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if _, ok := ctx.Get("userID"); !ok {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
		refreshCookie, err := ctx.Cookie("refresh_token")
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
		if err := c.service.LogoutUser(refreshCookie); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.SetCookie("access_token", "", -1, "/", "", false, true)
		ctx.SetCookie("refresh_token", "", -1, "/", "", false, true)
		ctx.JSON(http.StatusOK, gin.H{"message": "logged out"})
	}
}
