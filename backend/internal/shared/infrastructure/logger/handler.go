package logger

import (
	"context"
	"log/slog"
)

type ContextKey string

const (
	RequestKey ContextKey = "X-Request-ID"
)

type LogHandler struct {
	handler slog.Handler
}

func NewLogHandler(handler slog.Handler) *LogHandler {
	return &LogHandler{handler: handler}
}

func (h *LogHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &LogHandler{handler: h.handler.WithAttrs(attrs)}
}

func (h *LogHandler) WithGroup(name string) slog.Handler {
	return &LogHandler{handler: h.handler.WithGroup(name)}
}

func (h *LogHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.handler.Enabled(ctx, level)
}

func (h *LogHandler) Handle(ctx context.Context, record slog.Record) error {
	requestID, ok := ctx.Value(RequestKey).(string)
	if ok {
		record.AddAttrs(slog.String(string(RequestKey), requestID))
	}

	return h.handler.Handle(ctx, record)
}
