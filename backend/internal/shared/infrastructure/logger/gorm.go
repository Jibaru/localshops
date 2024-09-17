package logger

import (
	"context"
	"log/slog"
	"time"

	gormlogger "gorm.io/gorm/logger"
)

type GormLogger struct {
}

func NewGormLogger() *GormLogger {
	return &GormLogger{}
}

func (l *GormLogger) LogMode(gormlogger.LogLevel) gormlogger.Interface {
	return l
}

func (l *GormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	slog.InfoContext(ctx, msg, data...)
}

func (l *GormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	slog.WarnContext(ctx, msg, data...)
}

func (l *GormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	slog.ErrorContext(ctx, msg, data...)
}

func (l *GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin)
	sql, rows := fc()
	slog.InfoContext(ctx, "db trace", "sql", sql, "rows", rows, "elapsed", elapsed)
}
