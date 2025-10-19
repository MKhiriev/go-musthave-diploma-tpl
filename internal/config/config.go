package config

type DBConfig struct {
	DSN string
}

type Auth struct {
	PasswordHashKey string
}

type Config struct {
	ServerAddress string
	Auth          Auth
	DB            DBConfig
}

// TODO add env and cmd params
func GetConfigs() *Config {
	return &Config{
		ServerAddress: "localhost:8080",
		Auth:          Auth{PasswordHashKey: "hash"},
		DB:            DBConfig{DSN: "postgres://postgres:postgrespassword@localhost:5432/praktikum?sslmode=disable"},
	}
}
