package store

import (
	"go-musthave-diploma-tpl/internal/logger"

	"gorm.io/gorm"
)

type Database struct {
	UserRepository UserRepository
}

func NewDatabase(db *gorm.DB, logger *logger.Logger) (*Database, error) {
	return &Database{
		UserRepository: NewDBUserRepository(db, logger),
	}, nil
}
