package middleware

import (
	"fmt"
	"time"
	"wecat/common/app"
	"wecat/common/email"
	"wecat/common/errcode"
	"wecat/common/logger"
	"wecat/global"

	"github.com/gin-gonic/gin"
)

func Recovery() gin.HandlerFunc {

	defailtMailer := email.NewEmail(&email.SMTPInfo{
		Host:     global.EmailSetting.Host,
		Port:     global.EmailSetting.Port,
		IsSSL:    global.EmailSetting.IsSSL,
		UserName: global.EmailSetting.UserName,
		Password: global.EmailSetting.Password,
		From:     global.EmailSetting.From,
	})

	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logger.Errorf("panic recover err: %v", err)

				err := defailtMailer.SendMail(
					global.EmailSetting.To,
					fmt.Sprintf("异常抛出,发生时间: %d", time.Now().Unix()),
					fmt.Sprintf("错误信息: %v", err),
				)

				if err != nil {
					logger.Errorf("mail.SendMail err: %v", err)

				}

				app.NewResponse(c).ToErrorResponse(errcode.ServerError)
				c.Abort()
			}
		}()
		c.Next()
	}
}
