package logger

import (
	"github.com/gookit/slog"
	"github.com/gookit/slog/handler"
	"os"
)

type Logger struct {
	*slog.Logger
}

func New(level string) (*Logger, error) {
	newL := slog.New()

	err := os.MkdirAll("logs", 0744)
	if err != nil {
		return nil, err
	}

	file, err := os.OpenFile("logs/all.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0744)
	if err != nil {
		return nil, err
	}

	fileHandler := handler.NewHandler(file, slog.LevelByName(level))
	stdoutHandler := handler.NewHandler(os.Stdout, slog.LevelByName(level))

	newL.AddHandlers(fileHandler, stdoutHandler)
	return &Logger{newL}, nil
}
