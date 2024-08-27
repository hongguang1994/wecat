package middleware

import "github.com/gin-gonic/gin"

// 限流
func Sentinel() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
