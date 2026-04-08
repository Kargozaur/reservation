package user

import (
	"errors"
	"user-service/models"
	"user-service/schemas/request"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user request.RegisterSchema) (*models.User, error) {
	u := user.ToUserModel()
	if err := r.db.Create(u).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, errors.New("User already exists")
		}
		return nil, err
	}
	return u, nil
}

func (r *UserRepository) FindUserByEmail(email string) (*models.User, error) {
	var u models.User
	if err := r.db.Where("email = ?", email).First(&u).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("User not found")
		}
		return nil, err
	}
	return &u, nil
}

func (r *UserRepository) FindUserById(id uuid.UUID) (*models.User, error) {
	var u models.User
	if err := r.db.Where("id = ?", id).First(&u).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("User not found")
		}
		return nil, err
	}
	return &u, nil
}

func (r *UserRepository) UpdateUserName(id uuid.UUID, name request.UpdateNameSchema) (*models.User, error) {
	var u models.User
	if err := r.db.Where("id = ?", id).First(&u).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("User not found")
		}
		return nil, err
	}
	name.ToUserModel(&u)
	if err := r.db.Save(&u).Error; err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *UserRepository) UpdateUserEmail(id uuid.UUID, email request.UpdateEmailSchema) (*models.User, error) {
	var u models.User
	if err := r.db.Where("id = ?", id).First(&u).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("User not found")
		}
		return nil, err
	}
	email.ToUserModel(&u)
	if err := r.db.Save(&u).Error; err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *UserRepository) UpdateUserPassword(id uuid.UUID, password request.UpdatePasswordSchema) (*models.User, error) {
	var u models.User
	if err := r.db.Where("id = ?", id).First(&u).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("User not found")
		}
		return nil, err
	}
	password.ToUserModel(&u)
	if err := r.db.Save(&u).Error; err != nil {
		return nil, err
	}
	return &u, nil
}
