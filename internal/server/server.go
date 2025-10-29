package server

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/wb-go/wbf/ginext"
)

const (
	RWTimeout       = 15 * time.Second
	IdleTimeout     = 60 * time.Second
	ShutdownTimeout = 10 * time.Second
)

type server struct {
	engine *ginext.Engine
	addr   string
}

func New(addr string, engine *ginext.Engine) *server {
	return &server{
		engine: engine,
		addr:   addr,
	}
}

func (s *server) Run(ctx context.Context) error {
	httpServer := &http.Server{
		Addr:         s.addr,
		Handler:      s.engine,
		ReadTimeout:  RWTimeout,
		WriteTimeout: RWTimeout,
		IdleTimeout:  IdleTimeout,
	}

	srvErr := make(chan error, 1)
	go func() {
		log.Printf("HTTP сервер запущен на %s", s.addr)
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			srvErr <- err
		}
	}()

	select {
	case <-ctx.Done():
		log.Println("Получен сигнал о завршении работы, инициируем graceful shutdown...")

		shutdownCtx, cancel := context.WithTimeout(context.Background(), ShutdownTimeout)
		defer cancel()

		if err := httpServer.Shutdown(shutdownCtx); err != nil {
			log.Printf("Ошибка при завершении работы сервера: %v", err)
			return err
		}

		log.Println("HTTP сервер успешно остановлен")
		return nil

	case err := <-srvErr:
		log.Printf("HTTP сервер завершил работу с ошибкой: %v", err)
		return err
	}
}
