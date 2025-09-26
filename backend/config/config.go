package config

import (
	"os"
	"strconv"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Redis    RedisConfig
	JWT      JWTConfig
}

type ServerConfig struct {
	Port string
	Mode string
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

type RedisConfig struct {
	Addr     string
	Password string
	DB       int
}

type JWTConfig struct {
	Secret string
	Expire int
}

func Load() *Config {
	return &Config{
		Server: ServerConfig{
			Port: getEnv("SERVER_PORT", "8080"),
			Mode: getEnv("SERVER_MODE", "debug"),
		},
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "8.138.244.31"),
			Port:     getEnv("DB_PORT", "3306"),
			User:     getEnv("DB_USER", "fangwei"),
			Password: getEnv("DB_PASSWORD", "NMDYyLjF52KS5S6N"),
			Name:     getEnv("DB_NAME", "fangwei"),
		},
		Redis: RedisConfig{
			Addr:     getEnv("REDIS_ADDR", "8.138.244.31:6379"),
			Password: getEnv("REDIS_PASSWORD", "123456"),
			DB:       getEnvInt("REDIS_DB", 0),
		},
		JWT: JWTConfig{
			Secret: getEnv("JWT_SECRET", "anti-fake-system-secret"),
			Expire: getEnvInt("JWT_EXPIRE", 24),
		},
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}
