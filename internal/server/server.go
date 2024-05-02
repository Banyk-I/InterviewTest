package server

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"net/http"
)

type Server struct {
	server *http.Server
}

func NewServer(port string, router http.Handler) *Server {
	return &Server{
		server: &http.Server{
			Addr:    port,
			Handler: router,
		},
	}
}

func (s *Server) Run() error {
	fmt.Printf("Server run :%v", viper.GetString("port"))
	return s.server.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
