package main

import (
	"context"
	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
	"log"
	"monopoly-auth/configs"
	"monopoly-auth/internal/storage"
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
	cfg, err := configs.InitConfig()
	if err != nil {
		logger.Errorf("Error of initialization config file: %s", err.Error())
	}

	middleware.Logger = logger

	_, err = storage.NewPostgresDB(
		storage.DBConfig{
			Host:     cfg.Database.Host,
			Port:     cfg.Database.Port,
			Username: cfg.Database.Username,
			Password: cfg.Database.Password,
			DBName:   cfg.Database.DBName,
			SSLMode:  cfg.Database.SSLMode,
		})

	if err != nil {
		logger.Errorf("Error of initialization storage: %s", err.Error())
	}

	srv := router.NewServer(logger)
	go func() {
		host := cfg.Server.Host
		port := cfg.Server.Port
		if err = srv.Run(host, port, router.NewRouter()); err != nil {
			logger.Errorf("Error of starting the HTTP server: %s", err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	if err = srv.Shutdown(context.Background()); err != nil {
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
