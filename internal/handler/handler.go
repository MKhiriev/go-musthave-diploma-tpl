package handler

import (
	"go-musthave-diploma-tpl/internal/logger"
	"go-musthave-diploma-tpl/internal/service"
)

type Handler struct {
	authService       service.AuthService
	balanceService    service.BalanceService
	orderService      service.OrderService
	withdrawalService service.WithdrawalService

	logger *logger.Logger
}

func NewHandler(services *service.Services, logger *logger.Logger) Handler {
	logger.Info().Msg("handler created")
	return Handler{
		authService:       services.AuthService,
		balanceService:    services.BalanceService,
		orderService:      services.OrderService,
		withdrawalService: services.WithdrawalService,
		logger:            logger,
	}
}
