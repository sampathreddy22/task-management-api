package middleware

import (
	"bytes"
	"io"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ResponseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w *ResponseWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func ZapLogger(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// start timer
		start := time.Now()

		//Read request body
		var requestBody []byte
		if c.Request.Body != nil {
			requestBody, _ = io.ReadAll(c.Request.Body)
			c.Request.Body = io.NopCloser(bytes.NewBuffer(requestBody))
		}

		// Create a new response writer to capture the response body
		responseWriter := &ResponseWriter{
			ResponseWriter: c.Writer,
			body:           bytes.NewBuffer([]byte{}),
		}
		c.Writer = responseWriter

		// Process request
		c.Next()

		// Log request details
		latency := time.Since(start)
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// Log request details
		logger.Info("Request",
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("raw", raw),
			zap.String("request_body", string(requestBody)),
			zap.String("response_body", responseWriter.body.String()),
			zap.Duration("latency", latency),
		)
	}
}
