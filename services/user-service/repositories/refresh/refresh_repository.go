package refresh

import (
	"crypto/sha256"
	"user-service/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RefreshRepository struct {
	db *gorm.DB
}

func NewRefreshRepository(db *gorm.DB) *RefreshRepository {
	return &RefreshRepository{db: db}
}

func (r *RefreshRepository) SaveRefreshToken(refreshToken string, userId uuid.UUID) error {
	hash := sha256.New()
	hash.Write([]byte(refreshToken))
	return r.db.Create(models.RefreshToken{Token: string(hash.Sum(nil)), UserID: userId}).Error
}

func (r *RefreshRepository) GetRefreshToken(refreshToken string) (*models.RefreshToken, error) {
	hash := sha256.New()
	hash.Write([]byte(refreshToken))
	var token models.RefreshToken
	err := r.db.First(&token, "token = ?", string(hash.Sum(nil))).Error
	return &token, err
}

func (r *RefreshRepository) DeleteAllUserTokens(userId uuid.UUID) error {
	return r.db.Delete(&models.RefreshToken{}, "user_id = ?", userId).Error
}

func (r *RefreshRepository) DeleteToken(tokenString string) error {
	hash := sha256.New()
	hash.Write([]byte(tokenString))
	return r.db.Delete(&models.RefreshToken{}, "token = ?", string(hash.Sum(nil))).Error
}
