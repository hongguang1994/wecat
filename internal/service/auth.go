package service

import (
	"errors"
	"wecat/common/logger"
)

type AuthRequest struct {
	AppKey    string `form:"app_key" json:"app_key" binding:"required"`
	AppSecret string `form:"app_secret" json:"app_secret" binding:"required"`
}

func (svc *Service) CheckAuth(param *AuthRequest) error {
	logger.Infof("appkey: %s, appsecret: %s", param.AppKey, param.AppSecret)
	auth, err := svc.dao.GetAuth(param.AppKey, param.AppSecret)
	if err != nil {
		return err
	}
	if auth.ID > 0 {
		return nil
	}
	return errors.New("auth info does not exist")
}
