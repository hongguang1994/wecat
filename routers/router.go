package routers

import (
	"wecat/middleware"
	"wecat/middleware/jwt"
	v1 "wecat/routers/api/v1"

	"github.com/gin-gonic/gin"
)

var v1GETMap = map[string]gin.HandlerFunc{
	"/tags":         v1.GetTags,
	"/articles":     v1.GetArticles,
	"/articles/:id": v1.GetArticle,
}

var v1POSTMap = map[string]gin.HandlerFunc{
	"/tags":     v1.AddTag,
	"/articles": v1.AddArticle,
}

var v1PUTMap = map[string]gin.HandlerFunc{
	"/tags/:id":     v1.EditTag,
	"/articles/:id": v1.EditArticle,
}

var v1DELETEMap = map[string]gin.HandlerFunc{
	"/tags/:id":     v1.DeleteTag,
	"/articles/:id": v1.DeleteArticle,
}

func InitRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	middleware.InitMiddleware(r)

	apiv1 := r.Group("/api/v1", jwt.JWT())
	{
		for k, v := range v1GETMap {
			apiv1.GET(k, v)
		}

		for k, v := range v1POSTMap {
			apiv1.POST(k, v)
		}

		for k, v := range v1PUTMap {
			apiv1.PUT(k, v)
		}

		for k, v := range v1DELETEMap {
			apiv1.DELETE(k, v)
		}
	}

	return r
}
