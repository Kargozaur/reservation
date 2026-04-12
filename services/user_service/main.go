package main

import (
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"user-service/conf"
	"user-service/controller"
	"user-service/middleware"

	"github.com/gin-gonic/gin"
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

func main() {
	logger, file := NewLogger("logs/app.log")
	defer file.Close()
	r := gin.Default()
	r.Use(middleware.RequestTime(logger))
	r.Use(gin.Recovery())
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "default page",
		})
	})
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	db := conf.NewDBConf()
	api := r.Group("/api")
	controller.UserRouter(api, db, logger)
	r.Run(":8081")
}
