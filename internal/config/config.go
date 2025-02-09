package config

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	Database DatabaseConfig
	Server   ServerConfig
	Logging  LoggingConfig
}

type ServerConfig struct {
	Port string
	Host string
}

type LoggingConfig struct {
	Level  string
	Format string
}

func Load() (*Config, error) {

	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath("./config")
	v.AddConfigPath("../config")
	v.AddConfigPath("../../config")

	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}
	// Load environment variables from .env file

	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return &cfg, nil

}

func Validate(cfg *Config) error {
	if cfg.Database.DBUser == "" || cfg.Database.DBPassword == "" || cfg.Database.Host == "" || cfg.Database.Name == "" {
		return fmt.Errorf("database configuration is missing")
	}

	if cfg.Server.Port == "" {
		return fmt.Errorf("server port is required")
	}
	return nil
}
