package logger

import (
	"context"
	"log/slog"
)

type LoggerFactory interface {
	Logger(ctx context.Context) *slog.Logger
}

type LoggerFactorySlog struct{}

func (lf *LoggerFactorySlog) Logger(ctx context.Context) *slog.Logger {

	logger := slog.Default()
	return logger
}
