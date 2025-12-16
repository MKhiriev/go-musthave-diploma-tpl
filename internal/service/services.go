package service

import (
	"go-musthave-diploma-tpl/internal/adapter"
	"go-musthave-diploma-tpl/internal/config"
	"go-musthave-diploma-tpl/internal/logger"
	"go-musthave-diploma-tpl/internal/store"
)

type Services struct {
	AuthService
	BalanceService
	OrderService
	WithdrawalService
}

func NewServices(database *store.Database, adapters *adapter.Adapters, cfg *config.Auth, logger *logger.Logger) *Services {
	defer logger.Info().Msg("services are initialized")
	services := &Services{
		AuthService:       NewAuthService(database.UserRepository, database.UserBalanceRepository, cfg, logger),
		BalanceService:    NewBalanceService(database.BalanceRepository, logger),
		OrderService:      NewOrderService(database.OrderRepository, adapters.AccrualAdapter, logger),
		WithdrawalService: NewWithdrawalService(database.WithdrawalRepository, logger),
	}

	return services
}
