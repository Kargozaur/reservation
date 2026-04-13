package users

import (
	"user-service/auth/pass"
	"user-service/auth/token"
	"user-service/repositories/refresh"
	"user-service/repositories/user"
)

type UserService struct {
	jwtEncoder  token.IJWT
	userRepo    user.IUserRepository
	refreshRepo refresh.IRefreshRepository
	hasher      pass.IPwdHasher
}

func NewUserService(jwtEncoder token.IJWT, userRepo user.IUserRepository, refreshRepo refresh.IRefreshRepository, hasher pass.IPwdHasher) *UserService {
	return &UserService{
		jwtEncoder:  jwtEncoder,
		userRepo:    userRepo,
		refreshRepo: refreshRepo,
		hasher:      hasher,
	}
}
