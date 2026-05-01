package repository

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"tracking.xlkv.com/internal/domain"
)

type UserRepository struct {
	pool *pgxpool.Pool
}

func NewUserRepository(pool *pgxpool.Pool) *UserRepository {
	return &UserRepository{pool: pool}
}

func (r *UserRepository) Create(ctx context.Context, driver *domain.Driver) error {

	query := `
	INSERT INTO drivers(username,password,name)
	VALUES($1,$2,$3)
	RETURNING id, created_at;
	`

	err := r.pool.QueryRow(
		ctx, query,
		driver.UserName,
		driver.Password,
		driver.Name,
	).Scan(
		&driver.ID,
		&driver.CreatedAt,
	)

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return domain.ErrAlreadyExists
		}
		return err
	}

	return nil
}

func (r *UserRepository) GetByID(ctx context.Context, id int64) (*domain.Driver, error) {

	query := `
	SELECT id, name, username, created_at
	FROM drivers
	WHERE id = $1
	`

	driver := &domain.Driver{}

	err := r.pool.QueryRow(
		ctx, query,
		id,
	).Scan(
		&driver.ID,
		&driver.Name,
		&driver.UserName,
		&driver.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return driver, nil
}

func (r *UserRepository) GetByUsername(ctx context.Context, username string) (*domain.Driver, error) {

	query := `
	SELECT id, name, username, password,created_at,
	FROM drivers
	WHERE username = $1
	`

	driver := &domain.Driver{}

	err := r.pool.QueryRow(
		ctx, query,
		username,
	).Scan(
		&driver.ID,
		&driver.Name,
		&driver.UserName,
		&driver.Password,
		&driver.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return driver, nil
}
