package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func SetRequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := uuid.NewUUID()
		if err == nil {
			requestID := id.String()
			c.Request.Header.Set("X-Request-ID", requestID)
			c.Writer.Header().Set("X-Request-ID", requestID)
		} else {
			fmt.Println(err)
		}
	}
}
