package adapter

import (
	"go-musthave-diploma-tpl/internal/config"
	"go-musthave-diploma-tpl/internal/logger"
)

type Adapters struct {
	AccrualAdapter
}

func NewAdapters(cfg *config.Adapter, logger *logger.Logger) *Adapters {
	defer logger.Info().Msg("adapters are initialized")

	return &Adapters{
		AccrualAdapter: NewAccrualAdapter(cfg, logger),
	}
}
