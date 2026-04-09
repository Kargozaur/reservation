package users

import (
	"errors"
	"user-service/auth/pass"
	"user-service/auth/token"
	"user-service/repositories/refresh"
	"user-service/repositories/user"
	"user-service/schemas/request"
	"user-service/schemas/response"
	"user-service/validators/credential"
)

type UserService struct {
	phasher  pass.IPwdHasher
	jwt      token.IJWT
	repo     user.IUserRepository
	refresh  refresh.IRefreshRepository
	validate credential.IValidator
}

func NewUserService(phasher pass.IPwdHasher, jwt token.IJWT, repo user.IUserRepository, refresh refresh.IRefreshRepository, validator credential.IValidator) *UserService {
	return &UserService{
		phasher:  phasher,
		jwt:      jwt,
		repo:     repo,
		refresh:  refresh,
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
	if err := s.refresh.SaveRefreshToken(user.ID, refreshToken); err != nil {
		return "", "", err
	}
	return accessToken, refreshToken, nil
}

func (s *UserService) GetUser(tokenString string) (response.UserResponse, error) {
	id, err := s.jwt.GetUId(tokenString)
	if err != nil {
		return response.UserResponse{}, err
	}
	user, err := s.repo.FindUserById(id)
	if err != nil {
		return response.UserResponse{}, err
	}
	return response.ToUserResponse(user), nil
}

func (s *UserService) UpdateName(tokenString string, updateRequest request.UpdateNameSchema) error {
	id, err := s.jwt.GetUId(tokenString)
	if err != nil {
		return err
	}
	if err = s.repo.UpdateUserName(id, updateRequest); err != nil {
		return err
	}
	return nil
}

func (s *UserService) UpdateEmail(tokenString string, updateRequest request.UpdateEmailSchema) error {
	id, err := s.jwt.GetUId(tokenString)
	if err != nil {
		return err
	}
	if err = s.validate.ValidateEmail(updateRequest.Email); err != nil {
		return err
	}
	if err = s.repo.UpdateUserEmail(id, updateRequest); err != nil {
		return err
	}
	return nil
}

func (s *UserService) UpdatePassword(tokenString string, updateRequest request.UpdatePasswordSchema) error {
	id, err := s.jwt.GetUId(tokenString)
	if err != nil {
		return err
	}
	if errs := s.validate.ValidatePassword(updateRequest.Password); len(errs) > 0 {
		return errors.Join(errs...)
	}
	hash, err := s.phasher.Hash(updateRequest.Password)
	if err != nil {
		return err
	}
	updateRequest.SwapPassword(hash)
	if err = s.repo.UpdateUserPassword(id, updateRequest); err != nil {
		return err
	}
	return nil
}

func (s *UserService) LogoutUser(refreshToken string) error {
	if err := s.refresh.DeleteRefreshToken(refreshToken); err != nil {
		return err
	}
	return nil
}
