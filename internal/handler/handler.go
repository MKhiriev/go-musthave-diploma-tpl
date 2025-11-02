package handler

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog"
)

type Handlers struct {
	logger *zerolog.Logger
}

func NewHandler(logger *zerolog.Logger) Handlers {
	return Handlers{logger: logger}
}

func (h *Handlers) Init() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Recoverer)

	return router
}
