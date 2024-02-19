package main

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
	"log"
	"monopoly-auth"
	"monopoly-auth/router"
	"monopoly-auth/router/middleware"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	logger, err := initLogger()
	if err != nil {
		log.Fatalf("Error of initialization logger: %s", err.Error())
	}
	if err := initConfig(); err != nil {
		logger.Errorf("Error of initialization config file: %s", err.Error())
	}

	middleware.Logger = logger

	srv := monopoly_auth.NewServer(logger)
	go func() {
		host := viper.GetString("server.host")
		port := viper.GetString("server.port")
		if err := srv.Run(host, port, router.NewRouter()); err != nil {
			logger.Errorf("Error of starting the HTTP server: %s", err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	if err := srv.Shutdown(context.Background()); err != nil {
		logger.Errorf("Error of shuting down the HTTP server: %s", err.Error())
	}
}

func initLogger() (*logrus.Logger, error) {
	return &logrus.Logger{
		Out:   os.Stderr,
		Level: logrus.DebugLevel,
		Formatter: &prefixed.TextFormatter{
			DisableColors:   false,
			TimestampFormat: "2006-01-02 15:04:05",
			FullTimestamp:   true,
			ForceFormatting: true,
		},
	}, nil
}

func initConfig() error {
	viper.AddConfigPath("../../configs")
	viper.SetConfigName("default")
	return viper.ReadInConfig()
}
