package setting

import "github.com/spf13/viper"

type Setting struct {
	vp *viper.Viper
}

func NewSetting(config string) (*Setting, error) {
	vp := viper.New()
	if config == "" {
		vp.SetConfigName("config")
		vp.AddConfigPath("configs/")
		vp.SetConfigType("yaml")
	} else {
		vp.SetConfigFile(config)
	}

	if err := vp.ReadInConfig(); err != nil {
		return nil, err
	}

	return &Setting{vp}, nil
}
