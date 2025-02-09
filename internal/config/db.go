package config

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseConfig struct {
	Host            string
	Port            string
	Name            string
	DBUser          string `mapstructure:"DB_USER"`
	DBPassword      string `mapstructure:"DB_PASSWORD"`
	MaxConnections  int    `mapstructure:"max_connections"`
	IdleConnections int    `mapstructure:"idle_connections"`
}

func InitializeDatabase(cfg *Config) (*gorm.DB, error) {
	fmt.Printf("DB_USER: %s\n", os.Getenv("DB_USER"))
	fmt.Printf("DB_PASSWORD: %s\n", os.Getenv("DB_PASSWORD"))

	cfg.Database.DBUser = os.Getenv("DB_USER")
	cfg.Database.DBPassword = os.Getenv("DB_PASSWORD")
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cfg.Database.Host, cfg.Database.Port, cfg.Database.DBUser, cfg.Database.DBPassword, cfg.Database.Name)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	pgDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get database instance: %w", err)
	}

	pgDB.SetMaxOpenConns(cfg.Database.MaxConnections)
	pgDB.SetMaxIdleConns(cfg.Database.IdleConnections)
	return db, nil
}
