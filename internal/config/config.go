package config

import "time"

type DBConfig struct {
	DSN string
}

type Auth struct {
	PasswordHashKey string
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
		Auth:   Auth{PasswordHashKey: "hash"},
		Server: Server{ServerAddress: "localhost:8080", RequestTimeout: time.Duration(10) * time.Second},
		DB:     DBConfig{DSN: "postgres://postgres:postgrespassword@localhost:5432/praktikum?sslmode=disable"},
	}
}
