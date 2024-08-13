package routers

import (
	"wecat/logger"
	"wecat/middleware"

	"github.com/gin-gonic/gin"
)

var (
	routerNoCheckRole = make([]func(*gin.RouterGroup), 0)
	routerCheckRole   = make([]func(*gin.RouterGroup), 0)
)

func InitRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	middleware.InitMiddleware(r)

	apiv1 := r.Group("/api/v1")

	sysNoCheckRoleRouter(apiv1)

	sysCheckRoleRouter(apiv1)

	routes := r.Routes()
	for _, v := range routes {
		logger.Infof("method : %s path: %s", v.Method, v.Path)
	}

	return r
}

func sysCheckRoleRouter(r *gin.RouterGroup) {
	for _, f := range routerCheckRole {
		f(r)
	}
}

func sysNoCheckRoleRouter(r *gin.RouterGroup) {
	for _, f := range routerNoCheckRole {
		f(r)
	}
}
