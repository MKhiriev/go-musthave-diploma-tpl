package service

import (
	"context"
	"fmt"
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
	balance, err := b.balanceRepository.FindBalanceByUserId(ctx, userId)
	if err != nil {
		return models.Balance{}, fmt.Errorf("error occurred during getting user balance: %w", err)
	}

	return balance, nil
}
