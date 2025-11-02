package config

type Config struct {
	ServerAddress string
}

func GetConfigs() *Config {
	return &Config{ServerAddress: "localhost:8080"}
}
