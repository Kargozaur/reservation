package token

import "github.com/google/uuid"

type IJWT interface {
	VerifyToken(tokenString string) bool
	GetUId(tokenString string) (uuid.UUID, error)
}
