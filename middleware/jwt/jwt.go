package jwt

import (
	"net/http"
	"time"
	"wecat/pkg/util"

	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// var code int
		var data interface{}

		token := ctx.Query("token")
		if token == "" {
			// 没有token
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				// 认证错误
			} else if time.Now().Unix() > claims.ExpiresAt {
				// 过期
			}
		}

		if false {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 1,
				"msg":  "",
				"data": data,
			})

			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
