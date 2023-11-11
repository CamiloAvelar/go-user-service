package http

import (
	"context"
	"net/http"
	"time"

	infrainterfaces "github.com/CamiloAvelar/go-user-service/infrastructure/interfaces"
)

type Server interface {
	WaitShutdown(ctx context.Context)
}

type httpServer struct {
	Server *http.Server
}

func Get(i infrainterfaces.HttpServerInjections) *httpServer {
	return &httpServer{
		Server: &http.Server{
			Handler:      GetRouter(i),
			Addr:         i.Config.ServerPort,
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
		},
	}
}

func (s *httpServer) WaitShutdown(ctx context.Context) {
	s.Server.Shutdown(ctx)
}
