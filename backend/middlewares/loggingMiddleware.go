package middlewares

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		duration := time.Since(startTime)

		fmt.Printf("Request: %s %s, Duration: %v\n", c.Request.Method, c.Request.URL, duration)
	}
}
