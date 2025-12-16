package server

import (
	"go-musthave-diploma-tpl/internal/config"
	"net/http"
)

type Server struct {
	server *http.Server
}

func (s *Server) ServerRun(handler http.Handler, cfg *config.Server) error {
	s.server = &http.Server{
		Addr:    cfg.ServerAddress,
		Handler: handler,
	}

	return s.server.ListenAndServe()
}
