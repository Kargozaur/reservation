package refresh

import (
	"crypto/sha256"
	"encoding/hex"
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

func (r *RefreshRepository) SaveRefreshToken(userId uuid.UUID, refreshToken string) error {
	hash := sha256.Sum256([]byte(refreshToken))
	return r.db.Create(&models.RefreshToken{Token: hex.EncodeToString(hash[:]), UserID: userId}).Error
}

func (r *RefreshRepository) GetRefreshToken(refreshToken string) (*models.RefreshToken, error) {
	hash := sha256.Sum256([]byte(refreshToken))
	var token models.RefreshToken
	err := r.db.First(&token, "token = ?", hex.EncodeToString(hash[:])).Error
	return &token, err
}

func (r *RefreshRepository) DeleteAllUserTokens(userId uuid.UUID) error {
	return r.db.Delete(&models.RefreshToken{}, "user_id = ?", userId).Error
}

func (r *RefreshRepository) DeleteRefreshToken(tokenString string) error {
	hash := sha256.Sum256([]byte(tokenString))
	return r.db.Delete(&models.RefreshToken{}, "token = ?", hex.EncodeToString(hash[:])).Error
}
