package main

import (
	"log/slog"
	"os"
	"path/filepath"
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
}
