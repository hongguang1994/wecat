package errcode

var (
	ErrorGenerateCaptchaFail = NewError(30010001, "验证码生成失败")
	ErrorVerifyCaptchaFail   = NewError(30010002, "验证码验证失败")
)
