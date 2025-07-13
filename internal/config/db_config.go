package config

import "os"

// retrieve env or default values hardcoded
func GetEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// Database config - connects using app_user for runtime operations
func DatabaseConfig() (string, string, string, string, string) {
	host := GetEnv("DB_HOST", "postgres") // Changed default to postgres for containerized
	port := GetEnv("DB_PORT", "5432")
	user := GetEnv("APP_DB_USER", "app_user") // Use app_user for runtime
	password := GetEnv("APP_DB_PASSWORD", "app_password")
	name := GetEnv("DB_NAME", "monitoring_testing")
	return host, port, user, password, name
}

// Redis config
func RedisConfig() (string, string) {
	host := GetEnv("REDIS_HOST", "redis") // Changed default to redis for containerized
	port := GetEnv("REDIS_PORT", "6379")
	return host, port
}
