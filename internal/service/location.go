package service

import (
	"context"
	"log/slog"

	"tracking.xlkv.com/internal/domain"
)

type LocationRepository interface {
	Create(ctx context.Context, location *domain.Location) error
	GetByVehicleID(ctx context.Context, id int64, page, limit int) ([]domain.Location, error)
	GetByTripID(ctx context.Context, id int64, page, limit int) ([]domain.Location, error)
	GetByID(ctx context.Context, id int64) (*domain.Location, error)
}

type LocationCache interface {
	SetLocation(ctx context.Context, vehicleID int64, location domain.Location) error
	PublishLocation(ctx context.Context, vehicleID int64, location domain.Location) error
}

type LocationService struct {
	repo  LocationRepository
	cache LocationCache
}

func NewLocationService(repo LocationRepository, cache LocationCache) *LocationService {
	return &LocationService{
		repo:  repo,
		cache: cache,
	}
}

func (s *LocationService) Create(ctx context.Context, lat, lng float64, tripID int64) (*domain.Location, error) {

	location := domain.Location{
		Lat:    lat,
		Lng:    lng,
		TripID: tripID,
	}

	err := s.repo.Create(
		ctx,
		&location,
	)

	if err != nil {
		slog.Error("Location creating fail", "err", err)
		return nil, err
	}

	if err = s.cache.SetLocation(ctx, location.VehicleID, location); err != nil {
		slog.Error("redis set location fail", "err", err)
	}

	if err = s.cache.PublishLocation(ctx, location.VehicleID, location); err != nil {
		slog.Error("redis publish location fail", "err", err)
	}

	return &location, nil
}
