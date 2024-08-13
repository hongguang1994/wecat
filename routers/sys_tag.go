package routers

import (
	"wecat/middleware/jwt"
	v1 "wecat/routers/api/v1"

	"github.com/gin-gonic/gin"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerSysTagRouter)
}

func registerSysTagRouter(rg *gin.RouterGroup) {
	r := rg.Group("/tag").Use(jwt.JWT())
	{
		r.GET("/tags", v1.GetTags)
		r.POST("/tags", v1.AddTag)
		r.PUT("/tags/:id", v1.EditTag)
		r.DELETE("/articles/:id", v1.DeleteTag)
	}
}
