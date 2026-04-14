package user

import (
	"context"
	"user-service/models"
	"user-service/schemas"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(ctx context.Context, schema schemas.CreateUser) (schemas.UserResponse, error) {
	user := schema.ToModel()
	if err := r.db.WithContext(ctx).Create(&user).Error; err != nil {
		return schemas.UserResponse{}, err
	}
	return schemas.UserResponse{ID: user.ID, Email: user.Email, FirstName: user.FirstName, LastName: user.LastName}, nil
}

func (r *UserRepository) FindUserByID(ctx context.Context, id uuid.UUID) (schemas.UserResponse, error) {
	var user models.User
	if err := r.db.WithContext(ctx).First(&user, id).Error; err != nil {
		return schemas.UserResponse{}, err
	}
	return schemas.UserResponse{ID: user.ID, Email: user.Email, FirstName: user.FirstName, LastName: user.LastName}, nil
}

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	if err := r.db.WithContext(ctx).First(&user, "email = ?", email).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) UpdateUserName(ctx context.Context, id uuid.UUID, schema schemas.UpdateName) error {
	var user models.User
	if err := r.db.WithContext(ctx).First(&user, id).Error; err != nil {
		return err
	}
	schema.ToModel(&user)
	return r.db.WithContext(ctx).Save(&user).Error
}

func (r *UserRepository) UpdateUserEmail(ctx context.Context, id uuid.UUID, schema schemas.UpdateEmail) error {
	var user models.User
	if err := r.db.WithContext(ctx).First(&user, id).Error; err != nil {
		return err
	}
	schema.ToModel(&user)
	return r.db.WithContext(ctx).Save(&user).Error
}

func (r *UserRepository) UpdateUserPassword(ctx context.Context, id uuid.UUID, schema schemas.UpdatePassword) error {
	var user models.User
	if err := r.db.WithContext(ctx).First(&user, id).Error; err != nil {
		return err
	}
	schema.ToModel(&user)
	return r.db.WithContext(ctx).Save(&user).Error
}
