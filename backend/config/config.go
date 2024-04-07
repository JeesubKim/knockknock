package config

import (
	"fmt"
	"os"
)

type Config struct {
	Port       string
	DBUser     string
	DBPassword string
	DBAddress  string
	DBName     string
	JWTSecret  string
}

var Envs = initConfig()

func initConfig() Config {
	return Config{
		Port:       getEnv("PORT", "8080"),
		DBUser:     getEnv("DB_USER", "root"),
		DBPassword: getEnv("DB_PASSWORD", "knockknock_db_password"),
		// DBAddress:  fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.0.1"), getEnv("DB_PORT", "3306")),
		DBAddress: fmt.Sprintf("%s:%s", getEnv("DB_HOST", "192.18.137.35"), getEnv("DB_PORT", "3306")),
		DBName:    getEnv("DB_NAME", "knockknock"),
		JWTSecret: getEnv("JWT_SECRET", "knockknockrandomjwtsecretkey"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
