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

type DBUserRepository struct {
	logger *logger.Logger
	db     *gorm.DB
}

func NewDBUserRepository(db *gorm.DB, logger *logger.Logger) UserRepository {
	logger.Info().Msg("DBUserRepository created")
	return &DBUserRepository{
		db:     db,
		logger: logger,
	}
}

func (r *DBUserRepository) CreateUser(ctx context.Context, user models.User) error {
	err := r.db.WithContext(ctx).Create(&user).Error

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

	return err
}

func (r *DBUserRepository) FindUserByLogin(ctx context.Context, user models.User) (models.User, error) {
	var foundUser models.User
	err := r.db.WithContext(ctx).Find(&foundUser, "login", user.Login).Error

	var pgErr *pgconn.PgError
	// if postgres returns error
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		case pgerrcode.NoDataFound:
			return models.User{}, ErrUserNotFound
		default:
			return models.User{}, fmt.Errorf("unexpected DB error: %w", err)
		}
	}

	return foundUser, err
}
