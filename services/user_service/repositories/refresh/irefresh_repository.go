package refresh

import (
	"context"
	"user-service/models"

	"github.com/google/uuid"
)

type IRefreshRepository interface {
	SaveRefreshToken(ctx context.Context, userId uuid.UUID, token string) error
	GetRefreshToken(ctx context.Context, tokenString string) (*models.RefreshToken, error)
	DeleteRefreshToken(ctx context.Context, tokenString string) error
	DeleteAllUserTokens(ctx context.Context, userId uuid.UUID) error
}
