package db

import (
	"context"
	"fmt"
	"skello/internal/config"
	"skello/internal/logger"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// allows for connection reuse
var pool *pgxpool.Pool

func Init(ctx context.Context) error {
	host, port, user, password, dbName := config.DatabaseConfig()

	connString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=prefer",
		host, port, user, password, dbName)

	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return fmt.Errorf("failed to parse database config: %w", err)
	}

	// pool settings
	config.MaxConns = 25
	config.MinConns = 5
	config.MaxConnLifetime = time.Hour
	config.MaxConnIdleTime = time.Minute * 30
	config.HealthCheckPeriod = time.Minute * 5

	pool, err = pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return fmt.Errorf("failed to create connection pool: %w", err)
	}

	if err = pool.Ping(ctx); err != nil {
		pool.Close()
		return fmt.Errorf("failed to ping database: %w", err)
	}
	logger.Get().Info("Connected to PostgreSQL successfully")
	return nil
}

// MustInit initializes db connection pool and panics on error
// Useful for apps that need a db connection to operate
func MustInit(ctx context.Context) {
	if err := Init(ctx); err != nil {
		logger.Get().WithError(err).Fatal("Failed to initialize database connection")
	}
}

func Get() *pgxpool.Pool {
	return pool
}

func Close() {
	if pool != nil {
		pool.Close()
		logger.Get().Info("Database connection pool is closed")
	}
}

// TODO: refine
