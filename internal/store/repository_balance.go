package store

import (
	"context"
	"errors"
	"fmt"
	"go-musthave-diploma-tpl/internal/logger"
	"go-musthave-diploma-tpl/models"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
)

type balanceRepository struct {
	db     *gorm.DB
	logger *logger.Logger
}

func NewBalanceRepository(db *gorm.DB, logger *logger.Logger) BalanceRepository {
	logger.Debug().Msg("BalanceRepository created")
	return &balanceRepository{
		db:     db,
		logger: logger,
	}
}

func (b *balanceRepository) FindBalanceByUserId(ctx context.Context, userId int64) (models.Balance, error) {
	var balance models.Balance

	err := b.db.WithContext(ctx).
		Where("user_id = ?", userId).
		Find(&balance).Error

	var pgErr *pgconn.PgError
	// if postgres returns error
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		case pgerrcode.NoDataFound:
			return models.Balance{}, ErrNoBalanceFound
		default:
			return models.Balance{}, fmt.Errorf("unexpected DB error: %w", err)
		}
	}

	return balance, nil
}
