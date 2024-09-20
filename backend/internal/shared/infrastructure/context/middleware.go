package context

import (
	"github.com/gin-gonic/gin"
)

const (
	CorrelationIdHeader string = "X-Correlation-Id"
)

func SetCorrelationIDMiddleware(ctx *gin.Context) {
	if correlationId := ctx.GetHeader("X-Correlation-Id"); correlationId != "" {
		ctx.Set(string(CorrelationIdKey), correlationId)
	}

	ctx.Next()
}
