package service

import (
	"go-musthave-diploma-tpl/internal/config"
	"go-musthave-diploma-tpl/internal/logger"
	"go-musthave-diploma-tpl/internal/store"
)

type Services struct {
	AuthService
	BalanceService
	OrderService
}

func NewServices(database *store.Database, cfg *config.Auth, logger *logger.Logger) *Services {
	defer logger.Info().Msg("services are initialized")
	return &Services{
		AuthService:    NewAuthService(database.UserRepository, database.UserBalanceRepository, cfg, logger),
		BalanceService: NewBalanceService(database.BalanceRepository, logger),
		OrderService:   NewOrderService(database.OrderRepository, logger),
	}
}
