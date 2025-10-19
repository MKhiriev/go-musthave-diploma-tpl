package handler

import (
	"go-musthave-diploma-tpl/internal/logger"
	"go-musthave-diploma-tpl/internal/service"
	"net/http"
)

type Handler struct {
	logger      *logger.Logger
	authService service.AuthService
}

func (h *Handler) order(w http.ResponseWriter, r *http.Request) {
	// TODO: implement me!
}

func (h *Handler) getOrders(w http.ResponseWriter, r *http.Request) {
	// TODO: implement me!
}

func (h *Handler) getBalance(w http.ResponseWriter, r *http.Request) {
	// TODO: implement me!
}

func (h *Handler) withdraw(w http.ResponseWriter, r *http.Request) {
	// TODO: implement me!
}

func (h *Handler) getWithdrawals(w http.ResponseWriter, r *http.Request) {
	// TODO: implement me!
}

func NewHandler(services *service.Services, logger *logger.Logger) Handler {
	logger.Info().Msg("handler created")
	return Handler{
		authService: services.AuthService,
		logger:      logger,
	}
}
