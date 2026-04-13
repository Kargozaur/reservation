package user

import (
	"user-service/models"
	"user-service/schemas"

	"github.com/google/uuid"
)

type IUserRepository interface {
	CreateUser(schemas.CreateUser) (*models.User, error)
	UpdateUserName(id uuid.UUID, schema schemas.UpdateName) error
	UpdateUserEmail(id uuid.UUID, schema schemas.UpdateEmail) error
	UpdateUserPassword(id uuid.UUID, schema schemas.UpdatePassword) error
}
