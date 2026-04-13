package user

import (
	"context"
	"user-service/models"
	"user-service/schemas"

	"github.com/google/uuid"
)

type IUserRepository interface {
	CreateUser(schemas.CreateUser) (schemas.UserResponse, error)
	FindUserByID(ctx context.Context, id uuid.UUID) (schemas.UserResponse, error)
	UpdateUserName(ctx context.Context, id uuid.UUID, schema schemas.UpdateName) error
	UpdateUserEmail(ctx context.Context, id uuid.UUID, schema schemas.UpdateEmail) error
	UpdateUserPassword(ctx context.Context, id uuid.UUID, schema schemas.UpdatePassword) error
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
}
