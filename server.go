package todo

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

// Функция для запуска сервера
func (s * Server) Run(port string) error {
	s.httpServer = &http.Server {
		Addr: ":" + port,
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	return s.httpServer.ListenAndServe()
}

// Функция для отключения сервера
func (s * Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
