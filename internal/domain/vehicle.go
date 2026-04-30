package domain

import "time"

type Vehicle struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Number    int64  `json:"number"`
	Status    string `json:"status"`
	DriverID  int64  `json:"driver_id"`
	RouteID   int64  `json:"route_id"`
	CreatedAt time.Time  `json:"created_at"`
}
