package service

import (
	"go-musthave-diploma-tpl/internal/config"
	"go-musthave-diploma-tpl/internal/logger"
	"go-musthave-diploma-tpl/internal/store"
)

type Services struct {
	AuthService
	logger *logger.Logger
}

func NewServices(database *store.Database, cfg *config.Config, logger *logger.Logger) *Services {
	defer logger.Info().Msg("services are initialized")
	return &Services{
		AuthService: NewAuthService(database.UserRepository, &cfg.Auth, logger),
	}
}
