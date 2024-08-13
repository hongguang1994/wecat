package routers

import (
	"wecat/middleware/jwt"
	v1 "wecat/routers/api/v1"

	"github.com/gin-gonic/gin"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerSysArticleRouter)
}

func registerSysArticleRouter(rg *gin.RouterGroup) {
	r := rg.Group("/article").Use(jwt.JWT())

	{
		r.GET("/articles", v1.GetArticles)
		r.GET("/articles/:id", v1.GetArticle)
		r.POST("/articles", v1.AddArticle)
		r.PUT("/articles/:id", v1.EditArticle)
		r.DELETE("/articles/:id", v1.DeleteArticle)
	}

}
