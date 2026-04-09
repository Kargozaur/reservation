package refresh

import (
	"user-service/models"

	"github.com/google/uuid"
)

type IRefreshRepository interface {
	SaveRefreshToken(userId uuid.UUID, token string) error
	GetRefreshToken(tokenString string) (*models.RefreshToken, error)
	DeleteRefreshToken(tokenString string) error
	DeleteAllUserTokens(userId uuid.UUID) error
}
