package users

import (
	"errors"
	"user-service/auth/pass"
	"user-service/auth/token"
	"user-service/repositories/user"
	"user-service/schemas/request"
	"user-service/schemas/response"
	"user-service/validators/credential"

	"github.com/google/uuid"
)

type UserService struct {
	phasher  pass.IPwdHasher
	jwt      token.IJWT
	repo     user.IUserRepository
	validate credential.IValidator
}

func NewUserService(phasher pass.IPwdHasher, jwt token.IJWT, repo user.IUserRepository, validator credential.IValidator) *UserService {
	return &UserService{
		phasher:  phasher,
		jwt:      jwt,
		repo:     repo,
		validate: validator,
	}
}

func (s *UserService) RegisterUser(createRequest request.RegisterSchema) (response.UserResponse, error) {
	if err := s.validate.ValidateEmail(createRequest.Email); err != nil {
		return response.UserResponse{}, err
	}
	if errs := s.validate.ValidatePassword(createRequest.Password); len(errs) > 0 {
		return response.UserResponse{}, errors.Join(errs...)
	}
	hashedPass, err := s.phasher.Hash(createRequest.Password)
	if err != nil {
		return response.UserResponse{}, err
	}
	createRequest.SwapPassword(hashedPass)
	user, err := s.repo.CreateUser(createRequest)
	if err != nil {
		return response.UserResponse{}, err
	}
	return response.ToUserResponse(user), nil
}

func (s *UserService) LoginUser(loginRequest request.LoginSchema) (string, string, error) {
	if err := s.validate.ValidateEmail(loginRequest.Email); err != nil {
		return "", "", err
	}
	if errs := s.validate.ValidatePassword(loginRequest.Password); len(errs) > 0 {
		return "", "", errors.Join(errs...)
	}
	user, err := s.repo.FindUserByEmail(loginRequest.Email)
	if err != nil {
		return "", "", err
	}
	if err := s.phasher.VerifyPwd(loginRequest.Password, []byte(user.Password)); err != nil {
		return "", "", err
	}
	accessToken, err := s.jwt.CreateAccessToken(user.ID)
	if err != nil {
		return "", "", err
	}
	refreshToken, err := s.jwt.CreateRefreshToken(user.ID)
	if err != nil {
		return "", "", err
	}
	return accessToken, refreshToken, nil
}

func (s *UserService) GetUser(id uuid.UUID) (response.UserResponse, error) {
	user, err := s.repo.FindUserById(id)
	if err != nil {
		return response.UserResponse{}, err
	}
	return response.ToUserResponse(user), nil
}

func (s *UserService) UpdateName(id uuid.UUID, updateRequest request.UpdateNameSchema) (response.UserResponse, error) {
	user, err := s.repo.UpdateUserName(id, updateRequest)
	if err != nil {
		return response.UserResponse{}, err
	}
	return response.ToUserResponse(user), nil
}

func (s *UserService) UpdateEmail(id uuid.UUID, updateRequest request.UpdateEmailSchema) (response.UserResponse, error) {
	user, err := s.repo.UpdateUserEmail(id, updateRequest)
	if err != nil {
		return response.UserResponse{}, err
	}
	return response.ToUserResponse(user), nil
}

func (s *UserService) UpdatePassword(id uuid.UUID, updateRequest request.UpdatePasswordSchema) (response.UserResponse, error) {
	user, err := s.repo.UpdateUserPassword(id, updateRequest)
	if err != nil {
		return response.UserResponse{}, err
	}
	return response.ToUserResponse(user), nil
}
