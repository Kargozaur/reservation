package handlers

import (
	"context"
	"fmt"
	"log/slog"
	"strings"
	"user-service/generated"
	"user-service/schemas"
	"user-service/services/users"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GRPCHandler struct {
	generated.UnimplementedUserServiceServer
	service users.UserService
	logger  *slog.Logger
}

func NewGRPCHandler(service users.UserService, logger *slog.Logger) *GRPCHandler {
	return &GRPCHandler{service: service, logger: logger}
}

func (h *GRPCHandler) CreateUser(ctx context.Context, req *generated.RegisterData) (*generated.GetDataResponse, error) {
	if req == nil || req.UserData == nil {
		return nil, status.Error(codes.InvalidArgument, "user data is required")
	}
	schema := schemas.CreateUser{
		Email:     req.UserData.Email,
		Password:  req.UserData.Password,
		FirstName: req.UserData.FirstName,
		LastName:  req.UserData.LastName,
	}
	user, err := h.service.CreateUser(ctx, schema)
	if err != nil {
		h.writeLog(ctx, err)
		return nil, err
	}
	return user, nil
}

func (h *GRPCHandler) LoginUser(ctx context.Context, req *generated.UserData) (*generated.GetTokenResponse, error) {
	if req == nil || req.Email == "" || req.Password == "" {
		return nil, status.Error(codes.InvalidArgument, "user data is required")
	}
	schema := schemas.LoginUser{
		Email:    req.Email,
		Password: req.Password,
	}
	token, err := h.service.LoginUser(ctx, schema)
	if err != nil {
		h.writeLog(ctx, err)
		return nil, err
	}
	return token, nil
}

func (h *GRPCHandler) GetUser(ctx context.Context, req *generated.GetDataRequest) (*generated.GetDataResponse, error) {
	if req == nil || req.UserId == "" {
		return nil, status.Error(codes.InvalidArgument, "user data is required")
	}
	userID, err := uuid.Parse(req.UserId)
	if err != nil {
		h.writeLog(ctx, err)
		return nil, err
	}
	user, err := h.service.GetData(ctx, userID)
	if err != nil {
		h.writeLog(ctx, err)
		return nil, err
	}
	return user, nil
}

func (h *GRPCHandler) UpdateName(ctx context.Context, req *generated.UpdateNameRequest) (*generated.GetMessageResponse, error) {
	if req == nil || req.UserId == "" || (req.FirstName == nil && req.LastName == nil) {
		return nil, status.Error(codes.InvalidArgument, "user data is required")
	}
	userID, err := uuid.Parse(req.UserId)
	if err != nil {
		h.writeLog(ctx, err)
		return nil, status.Error(codes.Internal, err.Error())
	}
	schema := schemas.UpdateName{
		FirstName: req.FirstName,
		LastName:  req.LastName,
	}
	user := h.service.UpdateUserName(ctx, userID, schema)
	if !strings.Contains(user.Message.Message, "success") {
		return nil, status.Error(codes.Internal, user.Message.Message)
	}
	return user, nil
}

func (h *GRPCHandler) UpdatePassword(ctx context.Context, req *generated.UpdatePasswordRequest) (*generated.GetMessageResponse, error) {
	if req == nil || req.UserId == "" || req.Password == "" {
		return nil, status.Error(codes.InvalidArgument, "user data is required")
	}
	userID, err := uuid.Parse(req.UserId)
	if err != nil {
		h.writeLog(ctx, err)
		return nil, status.Error(codes.Internal, err.Error())
	}
	schema := schemas.UpdatePassword{
		Password: req.Password,
	}
	user := h.service.UpdatePassword(ctx, userID, schema)
	if !strings.Contains(user.Message.Message, "success") {
		return nil, status.Error(codes.Internal, user.Message.Message)
	}
	return user, nil
}

func (h *GRPCHandler) UpdateEmail(ctx context.Context, req *generated.UpdateEmailRequest) (*generated.GetMessageResponse, error) {
	if req == nil || req.UserId == "" || req.Email == "" {
		return nil, status.Error(codes.InvalidArgument, "user data is required")
	}
	userID, err := uuid.Parse(req.UserId)
	if err != nil {
		h.writeLog(ctx, err)
		return nil, status.Error(codes.Internal, err.Error())
	}
	schema := schemas.UpdateEmail{
		Email: req.Email,
	}
	user := h.service.UpdateEmail(ctx, userID, schema)
	if !strings.Contains(user.Message.Message, "success") {
		return nil, status.Error(codes.Internal, user.Message.Message)
	}
	return user, nil
}

func (h *GRPCHandler) RefreshToken(ctx context.Context, req *generated.GetTokenPair) (*generated.GetTokenResponse, error) {
	if req == nil || req.UserId == "" || req.Token == nil || req.Token.RefreshToken == "" {
		return nil, status.Error(codes.InvalidArgument, "user data is required")
	}
	userID, err := uuid.Parse(req.UserId)
	if err != nil {
		h.writeLog(ctx, err)
		return nil, status.Error(codes.Internal, err.Error())
	}
	token := h.service.RefreshToken(ctx, userID, req.Token.RefreshToken)
	return token, nil
}

func (h *GRPCHandler) writeLog(ctx context.Context, err error) {
	message := fmt.Sprintf("ERROR - USER_SERVICE - %s", err.Error())
	h.logger.ErrorContext(ctx, message)
}
