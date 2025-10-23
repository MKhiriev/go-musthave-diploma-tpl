package config

import "time"

type DBConfig struct {
	DSN string
}

type Auth struct {
	PasswordHashKey string
	TokenSignKey    string
	TokenIssuer     string
	TokenDuration   time.Duration
}

type Server struct {
	ServerAddress  string
	RequestTimeout time.Duration
}

type StructuredConfig struct {
	Auth   Auth
	DB     DBConfig
	Server Server
}

func GetStructuredConfig() (*StructuredConfig, error) {
	cfg, err := GetConfigs()
	if err != nil {
		return nil, err
	}

	return &StructuredConfig{
		Auth:   Auth{PasswordHashKey: cfg.HashKey, TokenSignKey: cfg.TokenSignKey, TokenIssuer: cfg.TokenIssuer, TokenDuration: time.Duration(cfg.TokenDuration) * time.Hour},
		Server: Server{ServerAddress: cfg.ServerAddress, RequestTimeout: time.Duration(cfg.RequestTimeout) * time.Second},
		DB:     DBConfig{DSN: cfg.DatabaseDSN},
	}, nil
}
