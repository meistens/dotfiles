package config

import "os"

// retrieve env or default values hardcoded
func GetEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// db config
func DatabaseConfig() (string, string, string, string, string) {
	host := GetEnv("DB_HOST", "localhost")
	port := GetEnv("DB_PORT", "5432")
	user := GetEnv("APP_DB_USER", "postgres")
	password := GetEnv("APP_DB_PASSWORD", "postgres")
	name := GetEnv("DB_NAME", "pzztgres")
	return host, port, user, password, name
}

// redis config
func RedisConfig() (string, string) {
	host := GetEnv("REDIS_HOST", "localhost")
	port := GetEnv("REDIS_PORT", "6379")
	return host, port
}

// TODO: refine
