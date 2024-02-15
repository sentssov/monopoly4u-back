package main

import (
	"context"
	"log"
	"monopoly-auth/router"
	"os"
	"os/signal"
	"syscall"
)

const (
	ADDR = "127.0.0.1"
	PORT = "8080"
)

func main() {
	srv := router.NewServer()
	go func() {
		if err := srv.Run(ADDR, PORT, router.NewHttpHandler()); err != nil {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Fatal(err)
	}
}
