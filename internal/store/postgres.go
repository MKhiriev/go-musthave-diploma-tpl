package store

import (
	"go-musthave-diploma-tpl/internal/config"
	"go-musthave-diploma-tpl/internal/logger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresConnection(cfg *config.DBConfig, log *logger.Logger) (*gorm.DB, error) {
	postgresDialector := postgres.Open(cfg.DSN)

	connection, err := gorm.Open(postgresDialector)
	if err != nil {
		log.Err(err).Msg("error during connecting to postgres db")
		return nil, err
	}

	log.Info().Msg("successful connection to db")
	return connection, nil
}
