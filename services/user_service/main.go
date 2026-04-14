package main

import (
	"context"
	"log"
	"log/slog"
	"net"
	"os"
	"path/filepath"
	"user-service/auth/pass"
	"user-service/auth/token"
	"user-service/conf"
	"user-service/generated"
	"user-service/handlers"
	"user-service/repositories/refresh"
	"user-service/repositories/user"
	"user-service/services/users"

	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func NewLogger(filePath string) (*slog.Logger, *os.File) {
	dir := filepath.Dir(filePath)
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		panic(err)
	}

	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	opts := &slog.HandlerOptions{
		Level:     slog.LevelInfo,
		AddSource: true,
	}

	handler := slog.NewJSONHandler(file, opts)
	return slog.New(handler), file
}

func buildDeps(db *gorm.DB, logger *slog.Logger) *handlers.GRPCHandler {
	jwtEncoder := token.NewJWT()
	hasher := pass.NewHasher(12)
	userRepo := user.NewUserRepository(db)
	refreshRepo := refresh.NewRefreshRepository(db)
	service := users.NewUserService(jwtEncoder, userRepo, refreshRepo, hasher)
	handler := handlers.NewGRPCHandler(*service, logger)
	return handler
}

func main() {
	logger, file := NewLogger("logs/user_service.log")
	defer file.Close()
	db := conf.NewDBConf()
	handler := buildDeps(db, logger)
	server := grpc.NewServer(grpc.UnaryInterceptor(loggingInterceptor(logger)))
	generated.RegisterUserServiceServer(server, handler)
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}
	server.Serve(lis)
}

func loggingInterceptor(logger *slog.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		logger.Info("gRPC request", "method", info.FullMethod)
		resp, err := handler(ctx, req)
		if err != nil {
			logger.Error("gRPC error", "method", info.FullMethod, "error", err)
		}
		return resp, err
	}
}
