package server

import (
	"context"
	"net/http"
	"time"
)

const (
	rTimeout       = 10 * time.Second
	wTimeout       = 10 * time.Second
	maxHeaderBytes = 1 << 20
)

type Server struct {
	httpServer *http.Server
}

func New(port string, handler http.Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:           ":" + port,
			Handler:        handler,
			ReadTimeout:    rTimeout,
			WriteTimeout:   wTimeout,
			MaxHeaderBytes: maxHeaderBytes,
		},
	}
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
