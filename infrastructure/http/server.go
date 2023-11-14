package http

import (
	"context"
	"net/http"
	"time"

	infrainterfaces "github.com/CamiloAvelar/go-user-service/infrastructure/interfaces"
)

type Server interface {
	WaitShutdown(ctx context.Context) error
}

type httpServer struct {
	http.Server
}

func GetServer(i *infrainterfaces.ServerInjections) *httpServer {
	return &httpServer{
		http.Server{
			Handler:      GetRouter(i),
			Addr:         i.Config.ServerPort,
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
		},
	}
}

func (s *httpServer) WaitShutdown(ctx context.Context) error {
	return s.Shutdown(ctx)
}
