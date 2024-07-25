package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "")
	})

	return r
}
