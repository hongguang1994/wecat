package middleware

import (
	"github.com/gin-gonic/gin"
)

// 权限检查中间件
func AuthCheckRole() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
