package main

import (
	"context"
	"github.com/joho/godotenv"
	"monopoly-auth/configs"
	"monopoly-auth/internal"
	"monopoly-auth/internal/storage"
	"monopoly-auth/pkg/logging"
	"monopoly-auth/router"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	logger := logging.InitLogger()
	cfg, err := configs.InitConfig()
	if err != nil {
		logger.Errorf("Error of initialization config file: %s", err.Error())
	}

	if err = godotenv.Load(); err != nil {
		logger.Errorf("Error of loading the .env file: %s", err.Error())
		return
	}

	db, err := storage.NewPostgresDB(
		storage.DBConfig{
			Host:     cfg.Database.Host,
			Port:     cfg.Database.Port,
			Username: cfg.Database.Username,
			DBName:   cfg.Database.DBName,
			SSLMode:  cfg.Database.SSLMode,
			Password: os.Getenv("POSTGRES_APP_DB_PASSWORD"),
		})

	if err != nil {
		logger.Errorf("Error of initialization storage: %s", err.Error())
	}

	m := internal.NewPlayerManager(db)
	h := router.NewHandler(m, logger)

	srv := router.NewServer(logger)
	go func() {
		host := cfg.Server.Host
		port := cfg.Server.Port
		if err = srv.Run(host, port, h.InitRoutes()); err != nil {
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
