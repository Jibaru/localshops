package logger

import (
	"log/slog"
	"time"

	"github.com/gin-gonic/gin"
)

func LogMiddleware(ctx *gin.Context) {
	start := time.Now()
	path := ctx.Request.URL.Path
	raw := ctx.Request.URL.RawQuery

	slog.InfoContext(
		ctx,
		"request started",
		slog.String("method", ctx.Request.Method),
		slog.String("path", path),
		slog.String("raw_query", raw),
	)

	// Process request
	ctx.Next()

	// Fill the params
	param := gin.LogFormatterParams{}

	param.TimeStamp = time.Now() // Stop timer
	param.Latency = param.TimeStamp.Sub(start)
	if param.Latency > time.Minute {
		param.Latency = param.Latency.Truncate(time.Second)
	}

	param.ClientIP = ctx.ClientIP()
	param.Method = ctx.Request.Method
	param.StatusCode = ctx.Writer.Status()
	param.ErrorMessage = ctx.Errors.ByType(gin.ErrorTypePrivate).String()
	param.BodySize = ctx.Writer.Size()
	if raw != "" {
		path = path + "?" + raw
	}
	param.Path = path

	slog.InfoContext(
		ctx,
		"request done",
		slog.String("client_ip", param.ClientIP),
		slog.String("method", param.Method),
		slog.Int("status", param.StatusCode),
		slog.String("path", param.Path),
		slog.String("latency", param.Latency.String()),
		slog.Int("body_size", param.BodySize),
		slog.String("error_message", param.ErrorMessage),
	)
}
