package main

import (
	"log/slog"
	"os"
	"path/filepath"
	"user-service/auth/pass"
	"user-service/auth/token"
	"user-service/handlers"
	"user-service/repositories/refresh"
	"user-service/repositories/user"
	"user-service/services/users"

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
}
