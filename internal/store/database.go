package store

import (
	"go-musthave-diploma-tpl/internal/logger"

	"gorm.io/gorm"
)

type Database struct {
	UserRepository    UserRepository
	BalanceRepository BalanceRepository
	OrderRepository   OrderRepository

	UserBalanceRepository UserBalanceRepository
}

func NewDatabase(db *gorm.DB, logger *logger.Logger) (*Database, error) {
	return &Database{
		UserRepository:        NewUserRepository(db, logger),
		BalanceRepository:     NewBalanceRepository(db, logger),
		OrderRepository:       NewOrderRepository(db, logger),
		UserBalanceRepository: NewUserBalanceRepository(db, logger),
	}, nil
}
