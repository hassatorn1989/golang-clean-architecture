package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	AppPort                string
	AppName                string
	AppEnv                 string
	DBHost                 string
	DBPort                 string
	DBUser                 string
	DBPassword             string
	DBName                 string
	DBParams               string
	JWTAccessSecret        string
	JWTRefreshSecret       string
	JWTIssuer              string
	JWTAccessExpireMinutes int
	JWTRefreshExpireDays   int
}

func Load() *Config {
	return &Config{
		AppPort:                getEnv("APP_PORT", "3000"),
		AppName:                getEnv("APP_NAME", "go-auth-rotation"),
		AppEnv:                 getEnv("APP_ENV", "development"),
		DBHost:                 getEnv("DB_HOST", "127.0.0.1"),
		DBPort:                 getEnv("DB_PORT", "3306"),
		DBUser:                 getEnv("DB_USER", "root"),
		DBPassword:             getEnv("DB_PASSWORD", "123456"),
		DBName:                 getEnv("DB_NAME", "go_auth_rotation"),
		DBParams:               getEnv("DB_PARAMS", "charset=utf8mb4&parseTime=True&loc=Local"),
		JWTAccessSecret:        getEnv("JWT_ACCESS_SECRET", "access-secret"),
		JWTRefreshSecret:       getEnv("JWT_REFRESH_SECRET", "refresh-secret"),
		JWTIssuer:              getEnv("JWT_ISSUER", "go-auth-rotation"),
		JWTAccessExpireMinutes: getEnvAsInt("JWT_ACCESS_EXPIRE_MINUTES", 15),
		JWTRefreshExpireDays:   getEnvAsInt("JWT_REFRESH_EXPIRE_DAYS", 7),
	}
}

func (c *Config) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", c.DBUser, c.DBPassword, c.DBHost, c.DBPort, c.DBName, c.DBParams)
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getEnvAsInt(key string, fallback int) int {
	valueStr := getEnv(key, "")
	if valueStr == "" {
		return fallback
	}
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return fallback
	}
	return value
}
