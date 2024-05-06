package logger

import (
	"context"
	"log/slog"
	"os"
)

const logFilePath = "cook_droogers.log"

type LoggerFactory interface {
	Logger(ctx context.Context) *slog.Logger
}

type LoggerFactorySlog struct{}

func (lf *LoggerFactorySlog) Logger(ctx context.Context) *slog.Logger {

	logfile, err := os.Create(logFilePath)
	if err != nil {
		return nil
	}

	logger := slog.New(slog.NewJSONHandler(logfile, &slog.HandlerOptions{Level: slog.LevelDebug}))

	return logger
}
