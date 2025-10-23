package service

import (
	"go-musthave-diploma-tpl/internal/config"
	"go-musthave-diploma-tpl/internal/logger"
	"go-musthave-diploma-tpl/internal/store"
)

type Services struct {
	AuthService
	BalanceService
}

func NewServices(database *store.Database, cfg *config.Auth, logger *logger.Logger) *Services {
	defer logger.Info().Msg("services are initialized")
	return &Services{
		AuthService:    NewAuthService(database.UserRepository, cfg, logger),
		BalanceService: NewBalanceService(database.BalanceRepository, logger),
	}
}
