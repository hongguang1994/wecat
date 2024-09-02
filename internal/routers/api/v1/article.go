package v1

import (
	"wecat/common/app"
	"wecat/common/errcode"

	"github.com/gin-gonic/gin"
)

type Aritcle struct{}

func NewArticle() Aritcle {
	return Aritcle{}
}

func (a Aritcle) Get(c *gin.Context) {
	app.NewResponse(c).ToErrorResponse(errcode.ServerError)
}

func (a Aritcle) List(c *gin.Context) {

}

func (a Aritcle) Create(c *gin.Context) {

}

func (a Aritcle) Update(c *gin.Context) {

}

func (a Aritcle) Delete(c *gin.Context) {

}
