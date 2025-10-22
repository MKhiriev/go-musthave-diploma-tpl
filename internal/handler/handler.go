package handler

import (
	"go-musthave-diploma-tpl/internal/logger"
	"go-musthave-diploma-tpl/internal/service"
)

type Handler struct {
	authService    service.AuthService
	balanceService service.BalanceService

	logger *logger.Logger
}

func NewHandler(services *service.Services, logger *logger.Logger) Handler {
	logger.Info().Msg("handler created")
	return Handler{
		authService:    services.AuthService,
		balanceService: services.BalanceService,
		logger:         logger,
	}
}
