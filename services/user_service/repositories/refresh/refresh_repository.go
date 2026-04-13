package refresh

import (
	"context"
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

func (r *RefreshRepository) SaveRefreshToken(ctx context.Context, userId uuid.UUID, refreshToken string) error {
	hash := sha256.Sum256([]byte(refreshToken))
	return r.db.WithContext(ctx).Create(&models.RefreshToken{Token: hex.EncodeToString(hash[:]), UserID: userId}).Error
}

func (r *RefreshRepository) GetRefreshToken(ctx context.Context, refreshToken string) (*models.RefreshToken, error) {
	hash := sha256.Sum256([]byte(refreshToken))
	var token models.RefreshToken
	err := r.db.WithContext(ctx).First(&token, "token = ?", hex.EncodeToString(hash[:])).Error
	return &token, err
}

func (r *RefreshRepository) DeleteAllUserTokens(ctx context.Context, userId uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&models.RefreshToken{}, "user_id = ?", userId).Error
}

func (r *RefreshRepository) DeleteRefreshToken(ctx context.Context, tokenString string) error {
	hash := sha256.Sum256([]byte(tokenString))
	return r.db.WithContext(ctx).Delete(&models.RefreshToken{}, "token = ?", hex.EncodeToString(hash[:])).Error
}
