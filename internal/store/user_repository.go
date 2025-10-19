package store

import (
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

func (r *DBUserRepository) CreateUser(user models.User) error {
	err := r.db.Create(user).Error

	var pgErr *pgconn.PgError
	// if postgres returns error
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		case pgerrcode.UniqueViolation:
			return LoginAlreadyExistsError
		default:
			return fmt.Errorf("unexpected DB error: %w", err)
		}
	}

	return err
}
