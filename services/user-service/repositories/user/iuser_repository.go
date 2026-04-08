package user

import (
	"user-service/schemas/request"
	"user-service/schemas/response"

	"github.com/google/uuid"
)

type IUserRepository interface {
	CreateUser(request.RegisterSchema) (response.UserResponse, error)
	FindUserByEmail(email string) (response.UserResponse, error)
	FindUserById(id uuid.UUID) (response.UserResponse, error)
	UpdateUserName(id uuid.UUID, name request.UpdateNameSchema) (response.UserResponse, error)
	UpdateUserEmail(id uuid.UUID, email request.UpdateEmailSchema) (response.UserResponse, error)
	UpdateUserPassword(id uuid.UUID, password request.UpdatePasswordSchema) (response.UserResponse, error)
}
