package store

import (
	"go-musthave-diploma-tpl/internal/logger"

	"gorm.io/gorm"
)

type Database struct {
	UserRepository    UserRepository
	BalanceRepository BalanceRepository
	OrderRepository   OrderRepository
}

func NewDatabase(db *gorm.DB, logger *logger.Logger) (*Database, error) {
	return &Database{
		UserRepository:    NewDBUserRepository(db, logger),
		BalanceRepository: NewBalanceRepository(db, logger),
		OrderRepository:   NewOrderRepository(db, logger),
	}, nil
}
