package refresh

import "github.com/google/uuid"

type IRefreshRepository interface {
	SaveRefreshToken(userId uuid.UUID, token string) error
	GetRefreshToken(userId uuid.UUID) (string, error)
	DeleteRefreshToken(tokenString string) error
	DeleteAllUserTokens(userId uuid.UUID) error
}
