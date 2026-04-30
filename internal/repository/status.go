package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type StatusRepositroy struct {
	pool *pgxpool.Pool
}

func NewStatusRepository(pool *pgxpool.Pool) *StatusRepositroy {
	return &StatusRepositroy{
		pool: pool,
	}
}

func (r StatusRepositroy) Status(ctx context.Context) (string, error) {
	return "ok", nil
}
