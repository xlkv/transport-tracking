package domain

import "time"

type Driver struct {
	ID        int64     `json:"id"`
	UserName  string    `json:"username"`
	Password  string    `json:"-"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}