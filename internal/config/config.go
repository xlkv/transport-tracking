package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Server   ServerConfig
	DB       DBConfig
	JWT      JWTConfig
	RedisUrl string
}

type ServerConfig struct {
	Port string
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

type JWTConfig struct {
	SecretKey    string
	ExpiresHours int
}

func (c *DBConfig) DSN() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.Host, c.Port, c.User, c.Password, c.DBName, c.SSLMode,
	)
}

func Load() (*Config, error) {
	godotenv.Load()

	cfg := &Config{
		Server: ServerConfig{
			Port: GetEnv("SERVER_PORT", "8000"),
		},
		DB: DBConfig{
			Host:    GetEnv("DB_HOST", "localhost"),
			Port:    GetEnv("DB_PORT", "5432"),
			User:    GetEnv("DB_USER", ""),
			Password:    GetEnv("DB_PASSWORD", ""),
			DBName:  GetEnv("DB_NAME", "transportdb"),
			SSLMode: GetEnv("DB_SSLMODE", "disable"),
		},
		JWT: JWTConfig{
			SecretKey:    GetEnv("SECRET_KEY", ""),
			ExpiresHours: 24,
		},
		RedisUrl: GetEnv("REDIS_URL", "redis://redis:6379"),
	}

	return cfg, nil
}

func GetEnv(key, defaultValue string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}

	return defaultValue
}
