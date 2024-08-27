package routers

import (
	"net/http"
	"time"
	"wecat/common/limiter"
	"wecat/global"
	"wecat/internal/middleware"
	v1 "wecat/internal/routers/api/v1"

	"github.com/gin-gonic/gin"
)

var methodLimiters = limiter.NewMethodLimiter().AddBuckets(limiter.LimiterBucketRule{
	Key:          "/auth",
	FillInterval: time.Second,
	Capacity:     10,
	Quantum:      10,
})

func NewRouter() *gin.Engine {
	r := gin.New()
	// r.Use(gin.Logger())
	// r.Use(gin.Recovery())
	// r.Use(middleware.Logger())
	r.Use(middleware.AccessLog())
	r.Use(middleware.Recovery())
	r.Use(middleware.Translations())

	r.Use(middleware.RateLimiter(methodLimiters))

	upload := v1.NewUpload()
	r.POST("/upload/file", upload.UploadFile)
	r.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))

	r.POST("/auth", v1.GetAuth)

	apiv1 := r.Group("api/v1")
	apiv1.Use(middleware.JWT())
	{
		tag := v1.NewTag()
		apiv1.POST("/rags", tag.Create)
		apiv1.DELETE("/tags/:id", tag.Delete)
		apiv1.PUT("/tags/:id", tag.Update)
		apiv1.PATCH("/tags/:id/state", tag.Update)
		apiv1.GET("/tags", tag.List)

		article := v1.NewArticle()
		apiv1.POST("/articles", article.Create)
		apiv1.DELETE("/articles/:id", article.Delete)
		apiv1.PUT("/articles/:id", article.Update)
		apiv1.PATCH("/articles/:id/state", article.Update)
		apiv1.GET("/articles/:id", article.Get)
		apiv1.GET("/articles", article.List)

	}

	return r
}
