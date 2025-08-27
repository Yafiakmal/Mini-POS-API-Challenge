package db

import (
	"os"
	"time"
	// "github.com/yafiakmal/Mini-POS-API-Challenge/"
)

type Config struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
	SSLMode  string
	TimeZone string
}

func GetDBConfig() *Config {
	return &Config{
		Host:     os.Getenv("DB_HOST"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		Port:     os.Getenv("DB_PORT"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		TimeZone: os.Getenv("DB_TIMEZONE"),
	}
}

type PoolConfig struct {
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime time.Duration
}

func GetPoolConfig() PoolConfig {
	return PoolConfig{
		MaxIdleConns:    5,
		MaxOpenConns:    20,
		ConnMaxLifetime: 30 * time.Minute,
	}
}
