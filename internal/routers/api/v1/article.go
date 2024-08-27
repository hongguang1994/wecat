package v1

import "github.com/gin-gonic/gin"

type Aritcle struct{}

func NewArticle() Aritcle {
	return Aritcle{}
}

func (a Aritcle) Get(c *gin.Context) {

}

func (a Aritcle) List(c *gin.Context) {

}

func (a Aritcle) Create(c *gin.Context) {

}

func (a Aritcle) Update(c *gin.Context) {

}

func (a Aritcle) Delete(c *gin.Context) {

}
