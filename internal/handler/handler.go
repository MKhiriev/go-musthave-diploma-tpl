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

func (h *Handlers) order(writer http.ResponseWriter, request *http.Request) {
	// TODO: implement me!
}

func (h *Handlers) getOrders(writer http.ResponseWriter, request *http.Request) {
	// TODO: implement me!
}

func (h *Handlers) getBalance(writer http.ResponseWriter, request *http.Request) {
	// TODO: implement me!
}

func (h *Handlers) withdraw(writer http.ResponseWriter, request *http.Request) {
	// TODO: implement me!
}

func (h *Handlers) getWithdrawals(writer http.ResponseWriter, request *http.Request) {
	// TODO: implement me!
}

func NewHandler(services *service.Services, logger *logger.Logger) Handlers {
	logger.Info().Msg("handler created")
	return Handlers{
		authService: services.AuthService,
		logger:      logger,
	}
}
