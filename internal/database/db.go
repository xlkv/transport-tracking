package database

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
	"tracking.xlkv.com/internal/config"
)

type DB struct {
	Pool *pgxpool.Pool
}

func New(cfg *config.DBConfig) (*DB, error) {

	pool, err := pgxpool.New(context.Background(), cfg.DSN())

	if err != nil {
		return nil, fmt.Errorf("db connection failed: %w", err)
	}

	if err := pool.Ping(context.Background()); err != nil {
		return nil, fmt.Errorf("db ping error %w", err)
	}

	slog.Info("DB connected")

	return &DB{Pool: pool}, nil
}
