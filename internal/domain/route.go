package domain

import "time"

type Route struct {
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	FromStopID int64     `json:"from_stop_id"`
	ToStopID   int64     `json:"to_stop_id"`
	CreatedAt  time.Time `json:"created_at"`
}
