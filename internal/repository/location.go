package repository

import (
	"context"
	"errors"
	"log/slog"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"tracking.xlkv.com/internal/domain"
)

type LocationRepository struct {
	pool *pgxpool.Pool
}

func NewLocationRepository(pool *pgxpool.Pool) *LocationRepository {
	return &LocationRepository{
		pool: pool,
	}
}

func (r *LocationRepository) Create(ctx context.Context, location *domain.Location) error {

	query := `
	INSERT INTO locations(lat, lng, vehicle_id, trip_id)
	SELECT $1, $2, t.vehicle_id, t.id
	FROM trips t
	WHERE id = $3
	RETURNING id,vehicle_id,recorded_at;
	`

	err := r.pool.QueryRow(
		ctx, query,
		location.Lat,
		location.Lng,
		location.TripID,
	).Scan(
		&location.ID,
		&location.VehicleID,
		&location.RecordedAt,
	)

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.ConstraintName {
			case "locations_vehicle_id_fkey":
				return domain.ErrInvalidParam
			case "locations_trip_id_fkey":
				return domain.ErrInvalidParam
			default:
				return err
			}
		}
		return err
	}

	return nil
}

func (r *LocationRepository) GetByVehicleID(ctx context.Context, id int64, page, limit int) ([]domain.Location, error) {

	offset := (page - 1) * limit

	query := `
	SELECT id, lat, lng, vehicle_id, trip_id, recorded_at
	FROM locations
	WHERE vehicle_id = $1
	ORDER BY recorded_at DESC
	LIMIT $2 OFFSET $3
	`

	rows, err := r.pool.Query(
		ctx, query,
		id, limit, offset,
	)

	if err != nil {
		return nil, domain.ErrNotFound
	}

	locations := make([]domain.Location, 0)

	for rows.Next() {
		location := domain.Location{}
		err := rows.Scan(
			&location.ID,
			&location.Lat,
			&location.Lng,
			&location.VehicleID,
			&location.TripID,
			&location.RecordedAt,
		)
		if err != nil {
			slog.Error("Get location by vehicle id", "err", err)
		} else {
			locations = append(locations, location)
		}
	}

	return locations, nil
}

func (r *LocationRepository) GetByTripID(ctx context.Context, id int64, page, limit int) ([]domain.Location, error) {

	offset := (page - 1) * limit

	query := `
	SELECT id, lat, lng, vehicle_id, trip_id, recorded_at
	FROM locations
	WHERE trip_id = $1
	ORDER BY recorded_at DESC
	LIMIT $2 OFFSET $3
	`

	rows, err := r.pool.Query(
		ctx, query,
		id, limit, offset,
	)

	if err != nil {
		return nil, domain.ErrNotFound
	}

	locations := make([]domain.Location, 0)

	for rows.Next() {
		location := domain.Location{}
		err := rows.Scan(
			&location.ID,
			&location.Lat,
			&location.Lng,
			&location.VehicleID,
			&location.TripID,
			&location.RecordedAt,
		)
		if err != nil {
			slog.Error("Get location by trip id", "err", err)
		} else {
			locations = append(locations, location)
		}
	}

	return locations, nil
}

func (r *LocationRepository) GetByID(ctx context.Context, id int64) (*domain.Location, error) {
	query := `
	SELECT id, lat, lng, vehicle_id, trip_id, recorded_at
	FROM locations
	WHERE id = $1
	`

	location := &domain.Location{}

	err := r.pool.QueryRow(
		ctx, query,
		id,
	).Scan(
		&location.ID,
		&location.Lat,
		&location.Lng,
		&location.VehicleID,
		&location.TripID,
		&location.RecordedAt,
	)

	if err != nil {
		return nil, domain.ErrNotFound
	}

	return location, nil
}
