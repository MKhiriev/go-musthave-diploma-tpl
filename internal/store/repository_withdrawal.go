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
	wr.logger.Debug().Int64("user_id", userId).Any("withdrawal before", withdrawal).Msg("[START]")

	result := wr.db.Debug().
		WithContext(ctx).Raw(withdrawSumWithBalanceCheck, map[string]interface{}{
		"order":   withdrawal.OrderNum,
		"sum":     withdrawal.Sum,
		"user_id": userId,
	}).Scan(&withdrawal)
	defer wr.logger.Debug().Int64("rows affected", result.RowsAffected).AnErr("error", result.Error).Int64("user_id", userId).Any("withdrawal after", withdrawal).Msg("[END]")

	var pgErr *pgconn.PgError
	if errors.As(result.Error, &pgErr) {
		switch pgErr.Code {
		case pgerrcode.UniqueViolation:
			wr.logger.Err(result.Error).Int64("user_id", userId).Str("return error", ErrWithdrawalForOrderAlreadyExists.Error()).Msg("CreateWithdrawal()")
			return ErrWithdrawalForOrderAlreadyExists
		case pgerrcode.NoData:
			wr.logger.Err(result.Error).Int64("user_id", userId).Str("return error", ErrWithdrawalWasNotCreated.Error()).Msg("CreateWithdrawal()")
			return ErrWithdrawalWasNotCreated
		default:
			wr.logger.Err(result.Error).Int64("user_id", userId).Msg("CreateWithdrawal(): unexpected DB error")
			return fmt.Errorf("unexpected DB error: %w", result.Error)
		}
	}

	if result.RowsAffected == 0 {
		wr.logger.Error().Err(result.Error).Int64("user_id", userId).Any("withdrawal", withdrawal).Str("return error", ErrWithdrawalWasNotCreated.Error()).Msg("CreateWithdrawal(): result row affected")
		return ErrWithdrawalWasNotCreated
	}

	return nil
}

func (wr *withdrawalRepository) GetWithdrawalsByUserId(ctx context.Context, userId int64) ([]models.Withdrawal, error) {
	var withdrawals []models.Withdrawal
	err := wr.db.Debug().WithContext(ctx).
		Where("user_id = ?", userId).
		Find(&withdrawals).Error

	if err != nil {
		return nil, fmt.Errorf("DB error: %w", err)
	}

	return withdrawals, nil
}
