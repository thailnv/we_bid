package middlewares

import (
	"bytes"
	"fmt"
	"io"
	"time"

	"github.com/gin-gonic/gin"
)

type bodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (bw bodyWriter) Write(b []byte) (int, error) {
	bw.body.Write(b)
	return bw.ResponseWriter.Write(b)
}

func SetupLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		requestID := c.Request.Header.Get("X-Request-ID")
		c.Request = c.Request.WithContext(c.Request.Context())
		buf, _ := io.ReadAll(c.Request.Body)
		requestBody := string(buf)
		fmt.Printf("Request: %s\nCall to: %s\nBody: %s\n", requestID, path, requestBody)
		cw := &bodyWriter{body: &bytes.Buffer{}, ResponseWriter: c.Writer}
		c.Writer = cw
		start := time.Now()
		c.Next()
		elapsed := time.Since(start)
		fmt.Printf("Request: %s\nComplete in: %s\nResponse: %s\n", requestID, elapsed, cw.body.String())
	}
}
