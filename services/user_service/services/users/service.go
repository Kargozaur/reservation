package users

import (
	"context"
	"errors"
	"user-service/auth/pass"
	"user-service/auth/token"
	pb "user-service/generated"
	"user-service/repositories/refresh"
	"user-service/repositories/user"
	"user-service/schemas"

	"github.com/google/uuid"
	"gorm.io/gorm"
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

func (s *UserService) CreateUser(ctx context.Context, schema schemas.CreateUser) (*pb.GetDataResponse, error) {
	hash, err := s.hasher.Hash(schema.Password)
	if err != nil {
		return nil, err
	}
	schema.SwapPassword(hash)
	user, err := s.userRepo.CreateUser(ctx, schema)
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, errors.New("email already exists")
		}
		return nil, err
	}
	return &pb.GetDataResponse{
		User: &pb.User{
			Id:        user.ID.String(),
			Email:     user.Email,
			FirstName: user.FirstName,
			LastName:  user.LastName,
		},
	}, nil
}

func (s *UserService) LoginUser(ctx context.Context, schema schemas.LoginUser) (*pb.GetTokenResponse, error) {
	user, err := s.userRepo.GetUserByEmail(ctx, schema.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("invalid credentials")
		}
		return nil, err
	}
	if err := s.hasher.Verify(user.Password, []byte(schema.Password)); err != nil {
		return nil, errors.New("invalid credentials")
	}
	accessToken, refreshToken, err := s.generateTokens(ctx, user.ID)
	if err != nil {
		return nil, err
	}
	return &pb.GetTokenResponse{
		Token: &pb.Token{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	}, nil
}

func (s *UserService) GetData(ctx context.Context, userId uuid.UUID) (*pb.GetDataResponse, error) {
	user, err := s.userRepo.FindUserByID(ctx, userId)
	if err != nil {
		return nil, err
	}
	return &pb.GetDataResponse{
		User: &pb.User{
			Id:        user.ID.String(),
			Email:     user.Email,
			FirstName: user.FirstName,
			LastName:  user.LastName,
		},
	}, nil
}

func (s *UserService) UpdateUserName(ctx context.Context, id uuid.UUID, schema schemas.UpdateName) *pb.GetMessageResponse {
	if err := s.userRepo.UpdateUserName(ctx, id, schema); err != nil {
		return &pb.GetMessageResponse{
			Message: &pb.DefaultResponse{
				Message: err.Error(),
			},
		}
	}
	return &pb.GetMessageResponse{
		Message: &pb.DefaultResponse{
			Message: "user name updated successfully",
		},
	}
}

func (s *UserService) UpdateEmail(ctx context.Context, id uuid.UUID, schema schemas.UpdateEmail) *pb.GetMessageResponse {
	if err := s.userRepo.UpdateUserEmail(ctx, id, schema); err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return &pb.GetMessageResponse{
				Message: &pb.DefaultResponse{
					Message: "email already exists",
				},
			}
		}
		return &pb.GetMessageResponse{
			Message: &pb.DefaultResponse{
				Message: err.Error(),
			},
		}
	}
	return &pb.GetMessageResponse{
		Message: &pb.DefaultResponse{
			Message: "email updated successfully",
		},
	}
}

func (s *UserService) UpdatePassword(ctx context.Context, id uuid.UUID, schema schemas.UpdatePassword) *pb.GetMessageResponse {
	hash, err := s.hasher.Hash(schema.Password)
	if err != nil {
		return &pb.GetMessageResponse{
			Message: &pb.DefaultResponse{
				Message: err.Error(),
			},
		}
	}
	schema.SwapPassword(hash)
	if err := s.userRepo.UpdateUserPassword(ctx, id, schema); err != nil {
		return &pb.GetMessageResponse{
			Message: &pb.DefaultResponse{
				Message: err.Error(),
			},
		}
	}
	return &pb.GetMessageResponse{
		Message: &pb.DefaultResponse{
			Message: "password updated successfully",
		},
	}
}

func (s *UserService) RefreshToken(ctx context.Context, userId uuid.UUID, tokenString string) *pb.GetTokenResponse {
	if err := s.refreshRepo.DeleteRefreshToken(ctx, tokenString); err != nil {
		return nil
	}

	accessToken, refreshToken, err := s.generateTokens(ctx, userId)
	if err != nil {
		return nil
	}

	return &pb.GetTokenResponse{
		Token: &pb.Token{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	}
}

func (s *UserService) generateTokens(ctx context.Context, userId uuid.UUID) (string, string, error) {
	accessToken, err := s.jwtEncoder.CreateAccessToken(userId)
	if err != nil {
		return "", "", err
	}
	refreshToken, err := s.jwtEncoder.CreateRefreshToken(userId)
	if err != nil {
		return "", "", err
	}
	if err := s.refreshRepo.SaveRefreshToken(ctx, userId, refreshToken); err != nil {
		return "", "", err
	}
	return accessToken, refreshToken, nil
}
