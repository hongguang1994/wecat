package server

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"math/big"
	"net/http"
	"os"
	"os/signal"
	"time"
	"wecat/logger"
	"wecat/routers"

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
	logger.Debug("Server run ...")

	router := routers.InitRouter()

	// server := &http3.Server{
	// 	Addr:       ":443",
	// 	Handler:    r,
	// 	TLSConfig:  http3.ConfigureTLSConfig(generateTLSConfig()),
	// 	QUICConfig: &quic.Config{},
	// }

	// if err := server.ListenAndServe(); err != nil {
	// 	logger.Fatal("faild to listen...")
	// }
	server := http.Server{
		Addr:    ":80",
		Handler: router,
		// TLSConfig: generateTLSConfig(),
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			logger.Fatal("faild to listen ...")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	logger.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		logger.Fatalf("Server Shutdown: %v", err)
	}

	logger.Info("Server exiting")
}

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
