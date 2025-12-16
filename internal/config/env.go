package config

import (
	"errors"
	"log"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	ServerAddress  string `env:"RUN_ADDRESS"`
	DatabaseDSN    string `env:"DATABASE_URI"`
	AccrualAddress string `env:"ACCRUAL_SYSTEM_ADDRESS"`

	HashKey      string `env:"KEY"`
	TokenSignKey string `env:"TOKEN_SIGN_KEY"`
	TokenIssuer  string `env:"TOKEN_ISSUER"`

	TokenDuration  int64 `env:"TOKEN_DURATION"`
	RequestTimeout int64 `env:"REQUEST_TIMEOUT"`
}

func GetConfigs() (*Config, error) {
	cfg := &Config{}
	err := env.Parse(cfg)
	if err != nil {
		log.Fatal(err)
	}

	// else get command line args or default values
	commandLineServerAddress, databaseDSN, commandLineAccrualAddress, commandLineTokenSignKey, commandLineHashKey,
		commandLineIssuer, commandLineTokenDuration, commandLineRequestTimeout := ParseConfigFlags()

	if cfg.ServerAddress == "" {
		cfg.ServerAddress = commandLineServerAddress
	}
	if cfg.DatabaseDSN == "" {
		cfg.DatabaseDSN = databaseDSN
	}
	if cfg.AccrualAddress == "" {
		cfg.AccrualAddress = commandLineAccrualAddress
	}
	if cfg.TokenSignKey == "" {
		cfg.TokenSignKey = commandLineTokenSignKey
	}
	if cfg.HashKey == "" {
		cfg.HashKey = commandLineHashKey
	}
	if cfg.TokenIssuer == "" {
		cfg.TokenIssuer = commandLineIssuer
	}
	if cfg.TokenDuration == 0 {
		cfg.TokenDuration = commandLineTokenDuration
	}
	if cfg.RequestTimeout == 0 {
		cfg.RequestTimeout = commandLineRequestTimeout
	}

	return cfg, cfg.Validate()
}

func (s *Config) Validate() error {
	switch {
	case s.ServerAddress == "":
		return errors.New("invalid Server Address")
	case s.DatabaseDSN == "":
		return errors.New("invalid Database DSN")
	}
	// TODO add accrual address check

	return nil
}
