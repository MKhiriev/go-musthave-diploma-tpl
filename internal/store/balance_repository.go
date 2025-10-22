package store

import (
	"context"
	"go-musthave-diploma-tpl/internal/logger"
	"go-musthave-diploma-tpl/models"

	"gorm.io/gorm"
)

type balanceRepository struct {
	db     *gorm.DB
	logger *logger.Logger
}

func NewBalanceRepository(db *gorm.DB, logger *logger.Logger) BalanceRepository {
	return &balanceRepository{
		db:     db,
		logger: logger,
	}
}

func (b *balanceRepository) FindBalanceByUserId(ctx context.Context, userId int64) (models.Balance, error) {
	//TODO implement me
	panic("implement me")
}
