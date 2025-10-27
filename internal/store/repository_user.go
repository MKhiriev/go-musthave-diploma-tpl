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

type userRepository struct {
	logger *logger.Logger
	db     *gorm.DB
}

func NewUserRepository(db *gorm.DB, logger *logger.Logger) UserRepository {
	logger.Debug().Msg("UserRepository created")
	return &userRepository{
		db:     db,
		logger: logger,
	}
}

func (r *userRepository) CreateUser(ctx context.Context, user models.User) error {
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

func (r *userRepository) FindUserByLogin(ctx context.Context, user models.User) (models.User, error) {
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
