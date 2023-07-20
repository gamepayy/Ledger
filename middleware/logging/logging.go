package logging

import (
	"bytes"
	"io"
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func GinzapLogger() gin.HandlerFunc {
	config := zap.Config{
		Level:            zap.NewAtomicLevelAt(zap.DebugLevel),
		OutputPaths:      []string{"logs/logs.json"},
		ErrorOutputPaths: []string{"logs/errors.json"},
		Encoding:         "json",
	}

	logger := zap.Must(config.Build())

	ginzapLogger := ginzap.GinzapWithConfig(logger, &ginzap.Config{
		TimeFormat: time.RFC3339,
		UTC:        true,
		TraceID:    true,
		Context: func(c *gin.Context) []zapcore.Field {
			var body []byte
			var buf bytes.Buffer
			if c.Request.Body != nil {
				tee := io.TeeReader(c.Request.Body, &buf)
				body, _ = io.ReadAll(tee)
				c.Request.Body = io.NopCloser(&buf)
			}

			return []zapcore.Field{
				zap.String("ip", c.ClientIP()),
				zap.String("method", c.Request.Method),
				zap.String("path", c.Request.URL.Path),
				zap.String("body", string(body)),
				zap.String("url", c.Request.URL.String()),
				zap.String("user-agent", c.Request.UserAgent()),
				zap.String("time", time.Now().Format(time.RFC3339)),
			}
		},
	})
	return ginzapLogger
}
