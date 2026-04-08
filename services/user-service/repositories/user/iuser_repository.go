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
	UpdateUserName(id uuid.UUID, name request.UpdateNameSchema) (*models.User, error)
	UpdateUserEmail(id uuid.UUID, email request.UpdateEmailSchema) (*models.User, error)
	UpdateUserPassword(id uuid.UUID, password request.UpdatePasswordSchema) (*models.User, error)
}
