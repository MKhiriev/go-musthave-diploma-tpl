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

type userBalanceRepository struct {
	logger *logger.Logger
	db     *gorm.DB
}

func NewUserBalanceRepository(db *gorm.DB, logger *logger.Logger) UserBalanceRepository {
	logger.Debug().Msg("UserBalanceRepository created")
	return &userBalanceRepository{
		db:     db,
		logger: logger,
	}
}

func (r *userBalanceRepository) CreateUserAndBalance(ctx context.Context, user models.User) (models.User, error) {
	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.Omit("user_id").Create(&user).Error
		var pgErr *pgconn.PgError
		// if postgres returns error
		if errors.As(err, &pgErr) {
			switch pgErr.Code {
			case pgerrcode.UniqueViolation:
				return ErrLoginAlreadyExists
			default:
				return fmt.Errorf("unexpected DB error: %w", err)
			}
		}

		err = tx.Create(&models.Balance{
			UserID: user.UserID,
		}).Error
		if err != nil {
			if errors.As(err, &pgErr) {
				switch pgErr.Code {
				case pgerrcode.UniqueViolation:
					return ErrBalanceAlreadyExists
				default:
					return fmt.Errorf("unexpected DB error: %w", err)
				}
			}
		}

		return nil
	})

	return user, err
}
