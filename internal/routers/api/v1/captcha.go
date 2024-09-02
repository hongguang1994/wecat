package v1

import (
	"wecat/common/app"
	"wecat/common/errcode"
	"wecat/common/logger"
	"wecat/common/util"

	"github.com/gin-gonic/gin"
)

var stringCaptcha = util.NewStringCaptcha()

func GenerateCaptcha(c *gin.Context) {
	response := app.NewResponse(c)
	id, b64s, answer, err := stringCaptcha.Generate()
	logger.Infof("captcha answer : %s", answer)
	if err != nil {
		response.ToErrorResponse(errcode.ErrorGenerateCaptchaFail)
	}

	response.ToOkResponse(gin.H{
		"captcha_id":   id,
		"captcha_data": b64s,
	})

}

func CaptchaVerify(c *gin.Context) {
	param := struct {
		CaptchaId     string `form:"captcha_id" json:"captcha_id" binding:"required"`
		CaptchaAnswer string `form:"captcha_answer" json:"captcha_answer" binding:"required"`
	}{}

	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		logger.Errorf("app.BindAndValid errs %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Error()))
		return
	}

	b := stringCaptcha.Verify(param.CaptchaId, param.CaptchaAnswer, true)
	if b {
		response.ToOkResponse(true)
	} else {
		response.ToErrorResponse(errcode.ErrorVerifyCaptchaFail)
	}

}
