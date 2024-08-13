package middleware

import (
	"wecat/middleware/logger"

	"github.com/gin-gonic/gin"
)

func InitMiddleware(r *gin.Engine) {
	r.Use(logger.Logger())

}
