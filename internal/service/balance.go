package service

import (
	"context"
	"go-musthave-diploma-tpl/internal/logger"
	"go-musthave-diploma-tpl/internal/store"
	"go-musthave-diploma-tpl/models"
)

type balanceService struct {
	balanceRepository store.BalanceRepository
	logger            *logger.Logger
}

func NewBalanceService(balanceRepository store.BalanceRepository, logger *logger.Logger) BalanceService {
	return &balanceService{
		balanceRepository: balanceRepository,
		logger:            logger,
	}
}

func (b *balanceService) GetBalanceByUserId(ctx context.Context, userId int64) (models.Balance, error) {
	//TODO implement me
	panic("implement me")
}
