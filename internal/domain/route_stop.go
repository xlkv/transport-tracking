package domain

import "time"

type RouteStop struct {
	ID        int64     `json:"id"`
	RouteID   int64     `json:"route_id"`
	StopID    int64     `json:"stop_id"`
	StopOrder int       `json:"stop_order"`
	CreatedAt time.Time `json:"created_at"`
}
