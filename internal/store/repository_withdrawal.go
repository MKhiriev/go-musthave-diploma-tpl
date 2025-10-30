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

type withdrawalRepository struct {
	db     *gorm.DB
	logger *logger.Logger
}

func NewWithdrawalRepository(db *gorm.DB, logger *logger.Logger) WithdrawalRepository {
	logger.Debug().Msg("WithdrawalRepository created")
	return &withdrawalRepository{
		db:     db,
		logger: logger,
	}
}

func (wr *withdrawalRepository) CreateWithdrawal(ctx context.Context, withdrawal models.Withdrawal, userId int64) error {
	result := wr.db.WithContext(ctx).
		Raw(withdrawSumWithBalanceCheck, map[string]interface{}{
			"order":   withdrawal.OrderNum,
			"sum":     withdrawal.Sum,
			"user_id": userId,
		}).Scan(&withdrawal)

	var pgErr *pgconn.PgError
	if errors.As(result.Error, &pgErr) {
		switch pgErr.Code {
		case pgerrcode.UniqueViolation:
			return ErrWithdrawalForOrderAlreadyExists
		case pgerrcode.NoData:
			return ErrWithdrawalWasNotCreated
		default:
			return fmt.Errorf("unexpected DB error: %w", result.Error)
		}
	}

	if result.RowsAffected == 0 {
		return ErrWithdrawalWasNotCreated
	}

	return nil
}

func (wr *withdrawalRepository) GetWithdrawalsByUserId(ctx context.Context, userId int64) ([]models.Withdrawal, error) {
	var withdrawals []models.Withdrawal
	err := wr.db.WithContext(ctx).
		Where("user_id = ?", userId).
		Find(&withdrawals).Error

	if err != nil {
		return nil, fmt.Errorf("DB error: %w", err)
	}

	return withdrawals, nil
}
