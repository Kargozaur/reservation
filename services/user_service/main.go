package main

import (
	"io"
	"log/slog"
	"net/http"
	"os"
	"user-service/conf"
	"user-service/controller"
	"user-service/middleware"

	"github.com/gin-gonic/gin"
)

func NewLogger(filePath string) (*slog.Logger, os.File) {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	opts := &slog.HandlerOptions{
		Level:     slog.LevelInfo,
		AddSource: true,
	}
	var mw io.Writer = file
	logger := slog.NewJSONHandler(mw, opts)
	return slog.New(logger), *file
}

func main() {
	logger, file := NewLogger("logs/app.log")
	defer file.Close()
	r := gin.Default()
	r.Use(middleware.RequestTime(logger))
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
