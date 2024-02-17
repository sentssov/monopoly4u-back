package main

import (
	"context"
	"github.com/spf13/viper"
	"log"
	"monopoly-auth"
	"monopoly-auth/router"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("Ошибка инициализации конфигурационного файла: %s", err.Error())
	}
	srv := monopoly_auth.NewServer()
	go func() {
		host := viper.GetString("server.host")
		port := viper.GetString("server.port")
		if err := srv.Run(host, port, router.NewRouter()); err != nil {
			log.Fatalf("Ошибка запуска HTTP-сервера: %s", err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Fatalf("Ошибка выключения сервера: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("default")
	return viper.ReadInConfig()
}
