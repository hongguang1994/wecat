package setting

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

type App struct {
	Host         string
	IsHttps      bool
	Port         string
	ReadTimeout  int
	WriteTimeout int
}

var AppSetting = &App{}

type Jwt struct {
	Secret  string
	Timeout int64
}

var JwtSetting = &Jwt{}

type Log struct {
	Compress      bool
	ConsoleStdout bool
	FileStdout    bool
	Level         string
	LocalTime     bool
	MaxAge        int
	MaxBackups    int
	MaxSize       int
	Path          string
}

var LogSetting = &Log{}

type SSL struct {
	Key string
	Pem string
}

var SSLSetting = &SSL{}

type Redis struct {
	Host        string
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout int
}

var RedisSetting = &Redis{}

func Setup(config string) error {

	viper.SetConfigFile(config)
	if err := viper.ReadInConfig(); err != nil {
		err = fmt.Errorf("parse Config file fail: %w", err)
		return err
	}

	appMap := viper.GetStringMap(`settings.application`)
	if err := mapstructure.Decode(appMap, AppSetting); err != nil {
		return err
	}

	jwtMap := viper.GetStringMap(`settings.jwt`)
	if err := mapstructure.Decode(jwtMap, JwtSetting); err != nil {
		return err
	}

	logMap := viper.GetStringMap(`settings.log`)
	if err := mapstructure.Decode(logMap, LogSetting); err != nil {
		return err
	}

	sslMap := viper.GetStringMap(`settings.ssl`)
	if err := mapstructure.Decode(sslMap, SSLSetting); err != nil {
		return err
	}

	return nil
}
