package token

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JWT struct {
	secret []byte
}

func NewJWT() *JWT {
	return &JWT{secret: []byte(os.Getenv("SECRET_KEY"))}
}

func (j *JWT) CreateAccessToken(userID uuid.UUID) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userID.String(),
		"iat": time.Now().UTC().Unix(),
		"exp": time.Now().UTC().Add(time.Minute * 15).Unix(),
		"iss": "user-service",
	})
	tokenString, err := token.SignedString(j.secret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (j *JWT) CreateRefreshToken(userID uuid.UUID) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userID.String(),
		"iat": time.Now().UTC().Unix(),
		"exp": time.Now().UTC().Add(time.Hour * 24 * 7).Unix(),
		"iss": "user-service",
	})
	tokenString, err := token.SignedString(j.secret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (j *JWT) VerifyToken(tokenString string) bool {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return j.secret, nil
	})
	return token.Valid && err == nil
}

func (j *JWT) GetId(tokenString string) (uuid.UUID, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return j.secret, nil
	})
	if err != nil {
		return uuid.Nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return uuid.Nil, jwt.ErrInvalidType
	}
	sub, ok := claims["sub"].(string)
	if !ok {
		return uuid.Nil, jwt.ErrInvalidType
	}
	return uuid.Parse(sub)
}
