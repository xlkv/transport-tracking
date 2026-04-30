package domain

import "time"

type Trip struct {
	ID        int64      `json:"id"`
	VehicleID int64      `json:"vehicle_id"`
	RouteID   int64      `json:"route_id"`
	StartedAt time.Time  `json:"started_at"`
	EndedAt   *time.Time `json:"ended_at"`
	Status    string     `json:"status"`
	CreatedAt time.Time  `json:"created_at"`
}
