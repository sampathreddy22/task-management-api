package config

import (
	"fmt"

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

	//Read YAML config file
	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	// Load .env file for secrets
	v.SetConfigName(".env")
	v.SetConfigType("env")
	v.AddConfigPath(".")
	v.AddConfigPath("../")
	v.AddConfigPath("../../")

	//Read .env file
	if err := v.MergeInConfig(); err != nil {
		return nil, fmt.Errorf("failed to merge .env file: %w", err)
	}

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return &cfg, nil

}

func Validate(cfg *Config) error {
	if cfg.Database.User == "" || cfg.Database.Password == "" || cfg.Database.Host == "" || cfg.Database.Name == "" {
		return fmt.Errorf("database configuration is missing")
	}

	if cfg.Server.Port == "" {
		return fmt.Errorf("server port is required")
	}
	return nil
}
