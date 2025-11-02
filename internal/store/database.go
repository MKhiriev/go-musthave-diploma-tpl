package store

import (
	"context"
	"fmt"
	"go-musthave-diploma-tpl/internal/logger"

	"gorm.io/gorm"
)

type Database struct {
	UserRepository        UserRepository
	BalanceRepository     BalanceRepository
	OrderRepository       OrderRepository
	WithdrawalRepository  WithdrawalRepository
	UserBalanceRepository UserBalanceRepository

	db     *gorm.DB
	logger *logger.Logger
}

func NewDatabase(db *gorm.DB, logger *logger.Logger) (*Database, error) {
	database := &Database{
		UserRepository:        NewUserRepository(db, logger),
		BalanceRepository:     NewBalanceRepository(db, logger),
		OrderRepository:       NewOrderRepository(db, logger),
		WithdrawalRepository:  NewWithdrawalRepository(db, logger),
		UserBalanceRepository: NewUserBalanceRepository(db, logger),
		db:                    db,
		logger:                logger,
	}

	err := database.Migrate(context.Background())
	if err != nil {
		return nil, fmt.Errorf("error occured during migration of DB: %w", err)
	}

	logger.Info().Msg("repositories are initialized")
	return database, nil
}

func (d *Database) Migrate(ctx context.Context) error {
	d.logger.Debug().Msg("creating tables if not exist")

	result := d.db.Debug().WithContext(ctx).Exec(createTablesIfNotExist)
	if err := result.Error; err != nil {
		d.logger.Err(err).Msg("error during creation of table occurred")
		return err
	}

	d.logger.Debug().Msg("tables are created successfully")
	return nil
}
