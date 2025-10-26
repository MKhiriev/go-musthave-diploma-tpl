package store

import (
	"go-musthave-diploma-tpl/internal/logger"

	"gorm.io/gorm"
)

type Database struct {
	UserRepository       UserRepository
	BalanceRepository    BalanceRepository
	OrderRepository      OrderRepository
	WithdrawalRepository WithdrawalRepository

	UserBalanceRepository UserBalanceRepository
}

func NewDatabase(db *gorm.DB, logger *logger.Logger) (*Database, error) {
	logger.Info().Msg("repositories are initialized")
	return &Database{
		UserRepository:        NewUserRepository(db, logger),
		BalanceRepository:     NewBalanceRepository(db, logger),
		OrderRepository:       NewOrderRepository(db, logger),
		WithdrawalRepository:  NewWithdrawalRepository(db, logger),
		UserBalanceRepository: NewUserBalanceRepository(db, logger),
	}, nil
}
