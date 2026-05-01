package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"tracking.xlkv.com/internal/domain"
)

type VehicleRepository struct {
	pool *pgxpool.Pool
}

func NewVehicleRepository(pool *pgxpool.Pool) *VehicleRepository {
	return &VehicleRepository{
		pool: pool,
	}
}

func (r *VehicleRepository) Create(ctx context.Context, vehicle *domain.Vehicle) error {
	query := `
	INSERT INTO vehicles(name,number,status,driver_id,route_id)
	VALUES($1,$2,$3,$4,$5)
	RETURNING id,created_at
	`

	return r.pool.QueryRow(
		ctx, query,
		vehicle.Name, vehicle.Number, vehicle.Status, vehicle.DriverID, vehicle.RouteID,
	).Scan(
		&vehicle.ID,
		&vehicle.CreatedAt,
	)
}

