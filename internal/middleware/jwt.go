package middleware

import (
	"wecat/common/app"
	"wecat/common/errcode"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// func JWT() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		// var code int
// 		var data interface{}

// 		token := ctx.Query("token")
// 		if token == "" {
// 			// 没有token
// 		} else {
// 			claims, err := util.ParseToken(token)
// 			if err != nil {
// 				// 认证错误
// 			} else if time.Now().Unix() > claims.ExpiresAt {
// 				// 过期
// 			}
// 		}

// 		if false {
// 			ctx.JSON(http.StatusUnauthorized, gin.H{
// 				"code": 1,
// 				"msg":  "",
// 				"data": data,
// 			})

// 			ctx.Abort()
// 			return
// 		}

// 		ctx.Next()
// 	}
// }

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			token string
			ecode = errcode.Success
		)

		if s, exist := c.GetQuery("token"); exist {
			token = s
		} else {
			token = c.GetHeader("token")
		}

		if token == "" {
			ecode = errcode.InvalidParams
		} else {
			_, err := app.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					ecode = errcode.UnauthorizedTokenTimeout
				default:
					ecode = errcode.UnauthorizedTokenError
				}
			}
		}

		if ecode != errcode.Success {
			response := app.NewResponse(c)
			response.ToErrorResponse(ecode)
			c.Abort()
			return
		}

		c.Next()

	}
}
