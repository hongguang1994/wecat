package server

import (
	"fmt"
	"wecat/logger"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	config string

	StartCmd = &cobra.Command{
		Use:     "server",
		Short:   "Start server",
		Example: "wecat server -c ./config/setting.yml",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return setup()
		},
		Run: func(cmd *cobra.Command, args []string) {
			run()
		},
	}
)

func init() {
	StartCmd.PersistentFlags().StringVarP(&config, "config", "c", "config/setting.yml", "configuration file.")
}

func setup() error {
	viper.SetConfigFile(config)
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Parse config file fail: %s", err.Error())
	}
	// 初始化日志
	logger.Init()

	return nil
}

func run() {
	logger.Debug("server run ...")

}
