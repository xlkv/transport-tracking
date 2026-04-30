package domain

import "time"

type Location struct {
	ID         int64     `json:"id"`
	Lat        float64   `json:"lat"`
	Lng        float64   `json:"lng"`
	VehicleID  int64     `json:"vehicle_id"`
	TripID     int64     `json:"trip_id"`
	RecordedAt time.Time `json:"recorded_at"`
}
