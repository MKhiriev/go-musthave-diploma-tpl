package service

import (
	"context"
	"errors"
	"fmt"
	"go-musthave-diploma-tpl/internal/logger"
	"go-musthave-diploma-tpl/internal/store"
	"go-musthave-diploma-tpl/internal/utils"
	"go-musthave-diploma-tpl/models"
)

type withdrawalService struct {
	withdrawalRepository store.WithdrawalRepository
	logger               *logger.Logger
}

func NewWithdrawalService(withdrawalRepository store.WithdrawalRepository, logger *logger.Logger) WithdrawalService {
	return &withdrawalService{
		withdrawalRepository: withdrawalRepository,
		logger:               logger,
	}
}

func (w *withdrawalService) Withdraw(ctx context.Context, withdrawal models.Withdrawal, userId int64) error {
	if withdrawal.Sum == 0 {
		return fmt.Errorf("invalid data was passed: %w", ErrWithdrawalSumIsZero)
	}

	isCorrect, err := utils.LunaCheckString(withdrawal.OrderNum)
	switch {
	case err != nil:
		return fmt.Errorf("%w: %w", ErrInvalidOrderNumber, err)
	case !isCorrect:
		return ErrNotCorrectOrderNumber
	}

	err = w.withdrawalRepository.CreateWithdrawal(ctx, withdrawal, userId)
	if err != nil {
		switch {
		case errors.Is(err, store.ErrWithdrawalWasNotCreated):
			return ErrInsufficientFunds
		//case errors.Is(err, store.ErrWithdrawalForOrderAlreadyExists):
		//	return ErrNotCorrectOrderNumber
		default:
			return fmt.Errorf("error during withdrawal: %w", err)
		}

	}

	return nil
}

func (w *withdrawalService) GetWithdrawals(ctx context.Context, userId int64) ([]models.Withdrawal, error) {
	withdrawals, err := w.withdrawalRepository.GetWithdrawalsByUserId(ctx, userId)
	if err != nil {
		return nil, fmt.Errorf("error occurred during getting order: %w", err)
	}
	if len(withdrawals) == 0 {
		return nil, ErrNoWithdrawalsFound
	}

	return withdrawals, nil
}
