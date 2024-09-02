package util

import (
	"github.com/mojocn/base64Captcha"
)

type Captcha struct {
	captcha *base64Captcha.Captcha
}

func NewStringCaptcha() *Captcha {
	store := base64Captcha.DefaultMemStore

	source := "1234567890abcdefghijkmnpqrstuvwxyzABCDEFGHIJKMNPQRSTUVWXYZ"

	driver := base64Captcha.NewDriverString(
		80,
		240,
		6,
		1,
		6,
		source,
		nil,
		nil,
		nil,
	)

	captcha := base64Captcha.NewCaptcha(driver, store)
	return &Captcha{
		captcha: captcha,
	}

}

func (c *Captcha) Generate() (id, b64s, answer string, err error) {
	id, b64s, answer, err = c.captcha.Generate()
	return id, b64s, answer, err
}

func (c *Captcha) Verify(id, answer string, clear bool) bool {
	return c.captcha.Verify(id, answer, clear)
}
