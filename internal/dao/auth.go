package dao

import "wecat/internal/models"

func (d *Dao) GetAuth(appKey, appSecret string) (models.Auth, error) {
	auth := models.Auth{AppKey: appKey, AppSecret: appSecret}
	return auth.Get(d.engine)
}
