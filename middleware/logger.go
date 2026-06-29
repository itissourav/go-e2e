package middleware

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Before request
		start := time.Now()
		path := c.Request.URL.Path

		// Process request
		c.Next()

		// After request
		duration := time.Since(start)
		fmt.Printf("[%s] %s %s in %v\n",
			c.Request.Method,
			path,
			strconv.Itoa(c.Writer.Status()),
			duration,
		)
	}
}
