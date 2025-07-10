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
	user := GetEnv("DB_USER", "myuser_change")
	password := GetEnv("DB_PASSWORD", "mypassword_change")
	name := GetEnv("DB_NAME", "mydb_change")
	return host, port, user, password, name
}

// redis config
func RedisConfig() (string, string) {
	host := GetEnv("REDIS_HOST", "localhost")
	port := GetEnv("REDIS_PORT", "6379")
	return host, port
}

// TODO: refine
