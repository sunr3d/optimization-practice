package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/sunr3d/optimization-practice/internal/config"
	"github.com/sunr3d/optimization-practice/internal/entrypoint"
)

func main() {
	log.Println("Запуск приложения...")

	log.Println("Загрузка конфигурации...")
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatalf("config.GetConfig: %v", err)
	}
	log.Printf("Конфигурация загружена: %+v", cfg)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	if err := entrypoint.Run(ctx, cfg); err != nil {
		log.Fatalf("entrypoint.Run: %v", err)
	}

	log.Println("Приложение остановлено")
}
