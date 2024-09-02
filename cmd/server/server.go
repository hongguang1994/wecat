package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"
	"wecat/common/logger"
	"wecat/common/setting"
	"wecat/global"

	"wecat/internal/model"
	"wecat/internal/routers"

	"github.com/spf13/cobra"
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
	StartCmd.PersistentFlags().StringVarP(&config, "config", "c", "configs/setting.yml", "configuration file.")
}

func setup() error {
	err := setupSetting()
	if err != nil {
		return err
	}

	err = setupLog()
	if err != nil {
		return err
	}

	err = setupDBEngine()
	if err != nil {
		return err
	}

	// err = setupRedis()
	// if err != nil {
	// 	return err
	// }

	return nil
}

func setupSetting() error {
	setting, err := setting.NewSetting(config)
	if err != nil {
		return err
	}

	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("JWT", &global.JWTSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("Email", &global.EmailSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("Log", &global.LogSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("Redis", &global.RedisSetting)
	if err != nil {
		return err
	}

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	global.JWTSetting.Expire *= time.Second

	return nil
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}

func setupLog() error {
	logger.Setup(global.LogSetting)
	return nil
}

func setupRedis() error {
	var err error
	global.RedisClient, err = global.NewRedisClient(global.RedisSetting)
	if err != nil {
		return err
	}
	return nil
}

func run() {
	logger.Debug("Server run ...")

	router := routers.NewRouter()

	// server := &http3.Server{
	// 	Addr:       ":443",
	// 	Handler:    r,
	// 	TLSConfig:  http3.ConfigureTLSConfig(generateTLSConfig()),
	// 	QUICConfig: &quic.Config{},
	// }

	server := http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	logger.Infof("listen: %s", ":"+global.ServerSetting.HttpPort)
	go func() {
		if global.ServerSetting.IsHttps {
			if err := server.ListenAndServeTLS(global.ServerSetting.SSL.Pem, global.ServerSetting.SSL.Key); err != nil && err != http.ErrServerClosed {
				logger.Fatal("faild to listen ...")
			}
		} else {
			if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				logger.Fatal("faild to listen ...")
			}
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	logger.Info("Shutdown Server ...")

	if err := server.Shutdown(ctx); err != nil {
		logger.Fatalf("Server Shutdown: %v", err)
	}

	logger.Info("Server exiting")
}

/*
func generateTLSConfig() *tls.Config {
	key, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		panic(err)
	}
	template := x509.Certificate{SerialNumber: big.NewInt(1)}
	certDER, err := x509.CreateCertificate(rand.Reader, &template, &template, &key.PublicKey, key)
	if err != nil {
		panic(err)
	}
	keyPEM := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(key)})
	certPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: certDER})

	tlsCert, err := tls.X509KeyPair(certPEM, keyPEM)
	if err != nil {
		panic(err)
	}
	return &tls.Config{
		Certificates: []tls.Certificate{tlsCert},
		NextProtos:   []string{"quic-echo-example"},
	}
}
*/
