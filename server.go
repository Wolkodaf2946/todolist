package todolist

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
	// будем использовать для запуска http сервера
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,          // максимальный размер заголовка (1 мб)
		ReadTimeout:    10 * time.Second, // таймаут на чтение
		WriteTimeout:   10 * time.Second, // таймаут на запись
	}

	return s.httpServer.ListenAndServe() // запуск сервера

}

// выход из приложения
func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
