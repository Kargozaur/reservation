package token

import "github.com/google/uuid"

type IJWT interface {
	CreateAccessToken(userID uuid.UUID) (string, error)
	CreateRefreshToken(userID uuid.UUID) (string, error)
}
