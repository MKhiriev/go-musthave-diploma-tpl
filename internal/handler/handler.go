package handler

import (
	"go-musthave-diploma-tpl/internal/logger"
	"go-musthave-diploma-tpl/internal/service"
	"net/http"
)

type Handlers struct {
	logger      *logger.Logger
	authService service.AuthService
}

func (h *Handlers) order(w http.ResponseWriter, r *http.Request) {
	// TODO: implement me!
}

func (h *Handlers) getOrders(w http.ResponseWriter, r *http.Request) {
	// TODO: implement me!
}

func (h *Handlers) getBalance(w http.ResponseWriter, r *http.Request) {
	// TODO: implement me!
}

func (h *Handlers) withdraw(w http.ResponseWriter, r *http.Request) {
	// TODO: implement me!
}

func (h *Handlers) getWithdrawals(w http.ResponseWriter, r *http.Request) {
	// TODO: implement me!
}

func NewHandler(services *service.Services, logger *logger.Logger) Handlers {
	logger.Info().Msg("handler created")
	return Handlers{
		authService: services.AuthService,
		logger:      logger,
	}
}
