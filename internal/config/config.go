package config

import "time"

type DBConfig struct {
	DSN string
}

type Auth struct {
	PasswordHashKey string
	TokenSignKey    string
	TokenIssuer     string
	TokenDuration   int64
}

type Server struct {
	ServerAddress  string
	RequestTimeout time.Duration
}

type Config struct {
	Auth   Auth
	DB     DBConfig
	Server Server
}

// TODO add env and cmd params
func GetConfigs() *Config {
	return &Config{
		Auth:   Auth{PasswordHashKey: "hash", TokenSignKey: "token", TokenIssuer: "gophermart", TokenDuration: 2},
		Server: Server{ServerAddress: "localhost:8080", RequestTimeout: time.Duration(10) * time.Second},
		DB:     DBConfig{DSN: "postgres://postgres:postgrespassword@localhost:5432/praktikum?sslmode=disable"},
	}
}
