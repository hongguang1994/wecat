package logger

import (
	"time"
	"wecat/logger"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		c.Next()

		endTime := time.Now()

		latencyTime := endTime.Sub(startTime)

		reqMethod := c.Request.Method

		reqUri := c.Request.RequestURI

		statusCode := c.Writer.Status()

		clientIp := c.ClientIP()

		logger.Infof("ts: %s statusCode: %3d latencyTime: %13v clientIp: %15s method: %s uri: %s",
			startTime.Format("2006-01-02 15:04:05.9999"),
			statusCode,
			latencyTime,
			clientIp,
			reqMethod,
			reqUri,
		)

	}
}
