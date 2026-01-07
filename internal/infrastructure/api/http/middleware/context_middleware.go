package middleware

import (
	"fmt"
	"io"
	"net/http"
	"runtime/debug"

	"github.com/InstayPMS/backend/internal/application/dto"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ContextMiddleware struct {
	log *zap.Logger
}

func NewContextMiddleware(log *zap.Logger) *ContextMiddleware {
	return &ContextMiddleware{log}
}

func (m *ContextMiddleware) Recovery() gin.HandlerFunc {
	return gin.CustomRecoveryWithWriter(io.Discard, func(c *gin.Context, err any) {
		var recErr error
		switch v := err.(type) {
		case error:
			recErr = v
		default:
			recErr = fmt.Errorf("%v", v)
		}

		stack := string(debug.Stack())
		m.log.Error("panic recovered",
			zap.String("path", c.Request.URL.Path),
			zap.String("method", c.Request.Method),
			zap.String("ip", c.ClientIP()),
			zap.String("user_agent", c.Request.UserAgent()),
			zap.String("stack", stack),
			zap.Error(recErr),
		)

		c.AbortWithStatusJSON(http.StatusInternalServerError, dto.APIResponse{
			InternalCode: 500000,
			Slug: "INTERNAL_SERVER_ERROR",
			Message: "internal server error",
		})
	})
}