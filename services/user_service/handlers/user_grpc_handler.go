package handlers

import (
	"context"
	"fmt"
	"log/slog"
	"user-service/generated"
	"user-service/schemas"
	"user-service/services/users"

	"github.com/google/uuid"
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

func (h *GRPCHandler) UpdateName(ctx context.Context, req *generated.UpdateNameRequest) *generated.GetMessageResponse {
	userID, err := uuid.Parse(req.UserId)
	if err != nil {
		h.writeLog(ctx, err)
		return nil
	}
	schema := schemas.UpdateName{
		FirstName: req.FirstName,
		LastName:  req.LastName,
	}
	user := h.service.UpdateUserName(ctx, userID, schema)
	return user
}

func (h *GRPCHandler) UpdatePassword(ctx context.Context, req *generated.UpdatePasswordRequest) *generated.GetMessageResponse {
	userID, err := uuid.Parse(req.UserId)
	if err != nil {
		h.writeLog(ctx, err)
		return nil
	}
	schema := schemas.UpdatePassword{
		Password: req.Password,
	}
	user := h.service.UpdatePassword(ctx, userID, schema)
	return user
}

func (h *GRPCHandler) UpdateEmail(ctx context.Context, req *generated.UpdateEmailRequest) *generated.GetMessageResponse {
	userID, err := uuid.Parse(req.UserId)
	if err != nil {
		h.writeLog(ctx, err)
		return nil
	}
	schema := schemas.UpdateEmail{
		Email: req.Email,
	}
	user := h.service.UpdateEmail(ctx, userID, schema)
	return user
}

func (h *GRPCHandler) RefreshToken(ctx context.Context, req *generated.GetTokenPair) *generated.GetTokenResponse {
	userID, err := uuid.Parse(req.UserId)
	if err != nil {
		h.writeLog(ctx, err)
		return nil
	}
	token := h.service.RefreshToken(ctx, userID, req.Token.RefreshToken)
	return token
}

func (h *GRPCHandler) writeLog(ctx context.Context, err error) {
	message := fmt.Sprintf("ERROR - USER_SERVICE - %s", err.Error())
	h.logger.ErrorContext(ctx, message)
}
