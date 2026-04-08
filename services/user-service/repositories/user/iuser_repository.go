package user

import (
	"user-service/models"
	"user-service/schemas/request"

	"github.com/google/uuid"
)

type IUserRepository interface {
	CreateUser(request.RegisterSchema) (*models.User, error)
	FindUserByEmail(email string) (*models.User, error)
	FindUserById(id uuid.UUID) (*models.User, error)
	UpdateUserName(id uuid.UUID, name request.UpdateNameSchema) error
	UpdateUserEmail(id uuid.UUID, email request.UpdateEmailSchema) error
	UpdateUserPassword(id uuid.UUID, password request.UpdatePasswordSchema) error
}
