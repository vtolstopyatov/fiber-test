package config

import (
	"os"
	"strconv"
	"time"
	// "string"
)

type PostgresConfig struct {
	User               string
	Password           string
	Name               string
	Host               string
	Port               string
	SSLMode            string
	MaxOpenConns       int
	MaxIdleConns       int
	ConnMaxLifetime time.Duration
}

type Config struct {
	Postgres    PostgresConfig
	Environment string
	HTTPPort	int
}

func GetConfig() *Config {
	return &Config{
		Postgres: PostgresConfig{
			User:               getEnv("POSTGRES_USER", "postgres"),
			Password:           getEnv("POSTGRES_PASSWORD", "postgres"),
			Name:               getEnv("DB_NAME", "postgres"),
			Host:               getEnv("DB_HOST", "localhost"),
			Port:               getEnv("DB_PORT", "5432"),
			SSLMode:            getEnv("DB_SSLMODE", "disable"),
			MaxOpenConns:       getEnvAsInt("DB_MAX_OPEN_CONNS", 10),
			MaxIdleConns:       getEnvAsInt("DB_MAX_IDLE_CONNS", 10),
			ConnMaxLifetime: time.Duration(getEnvAsInt("DB_CONN_MAX_LIFETIME", 5)),
		},
		Environment: getEnv("ENVIRONMENT", "development"),
		HTTPPort: getEnvAsInt("HTTP_PORT", 3000),
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultVal
}
