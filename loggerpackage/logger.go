package loggerpackage

import (
	"context"
	"fmt"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger holds the structured logger
type Logger struct {
	log          *zap.Logger
	functionLogs sync.Map // Using sync.Map for lock-free access
}

// ContextKey prevents collisions in context
type ContextKey string

const loggerKey ContextKey = "logger"

// const functionKey ContextKey = "function"

// New creates a new logger instance
func New(level zapcore.Level) *Logger {
	config := zap.NewProductionConfig()
	config.Level = zap.NewAtomicLevelAt(level)
	logger, _ := config.Build()

	return &Logger{
		log: logger,
	}
}

// WithContext attaches the logger to a context
func (l *Logger) WithContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, loggerKey, l)
}

// SetFunctionLevel allows dynamic log levels for functions
func (l *Logger) SetFunctionLevel(ctx context.Context, function string, level zapcore.Level) {
	l.functionLogs.Store(function, level)
}

// AddMetadata attaches metadata to the logger
func (l *Logger) AddMetadata(ctx context.Context, key string, value interface{}) context.Context {
	return context.WithValue(ctx, ContextKey(key), value)
}

// WithFunctionContext returns a function-specific logger
func WithFunctionContext(ctx context.Context, key string) (*zap.SugaredLogger, error) {
	logger, ok := ctx.Value(loggerKey).(*Logger)
	if !ok {
		return nil, fmt.Errorf("logger not found in context")
	}

	// Retrieve the function-specific log level, if it exists
	if level, exists := logger.functionLogs.Load(key); exists {
		// Create a new logger with the updated level
		config := zap.NewProductionConfig()
		config.Level = zap.NewAtomicLevelAt(level.(zapcore.Level))
		newLogger, _ := config.Build()
		logger.log = newLogger
	}

	// Create the sugared logger with the function context
	sugaredLogger := logger.log.Sugar().With("function", key)

	// Add any additional metadata from context if available
	if key != "" {
		if parent, ok := ctx.Value(ContextKey(key)).(map[string]interface{}); ok {
			for k, val := range parent {
				sugaredLogger = sugaredLogger.With(k, val)
			}
		}
	}

	return sugaredLogger, nil
}
