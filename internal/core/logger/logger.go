package core_logger

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"time"
)

type Logger struct {
	*slog.Logger
	file *os.File
}

type loggerContext struct{}

var defaultLogger = &Logger{
	Logger: slog.Default(),
}

var (
	loggerContextKey = loggerContext{}
)

func NewLogger(config *Config) (*Logger, error) {
	if err := os.MkdirAll(config.LogFolder, 0o755); err != nil {
		return nil, fmt.Errorf("failed to create log folder: %w", err)
	}

	timestamp := time.Now().UTC().Format("2006-01-02T15-04-05.000000")
	logFilePath := filepath.Join(config.LogFolder, timestamp+".log")

	logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o644)
	if err != nil {
		return nil, fmt.Errorf("open log file: %w", err)
	}

	writer := io.MultiWriter(os.Stdout, logFile)

	opts := &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.Level(config.LogLevel),
	}

	loggerHandler := slog.NewTextHandler(writer, opts)
	logger := slog.New(loggerHandler)

	return &Logger{
		Logger: logger,
		file:   logFile,
	}, nil
}

func (l *Logger) With(fields ...any) *Logger {
	return &Logger{
		Logger: l.Logger.With(fields...),
		file:   l.file,
	}
}

func (l *Logger) Close() error {
	if l.file != nil {
		return l.file.Close()
	}
	return nil
}

func WithLogger(ctx context.Context, logger *Logger) context.Context {
	return context.WithValue(ctx, loggerContextKey, logger)
}

func FromContext(ctx context.Context) *Logger {
	l, ok := ctx.Value(loggerContextKey).(*Logger)
	if !ok {
		return defaultLogger
	}

	return l
}
