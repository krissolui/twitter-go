package server

import (
	"auth-service/internal/config"
	"fmt"
	"log"
	"net/http"
)

type Server struct {
	config  *config.Config
	handler http.Handler
	http    *http.Server
}

func NewServer(config *config.Config, handler http.Handler) Server {
	return Server{
		config:  config,
		handler: handler,
		http: &http.Server{
			Addr:    fmt.Sprintf(":%s", config.Port),
			Handler: handler,
		},
	}
}

func (s *Server) Start() error {
	log.Printf("Starting server on port %s\n", s.config.Port)
	return s.http.ListenAndServe()
}

func (s *Server) Stop() {
	s.http.Close()
}
