package domain

import "time"

type Driver struct {
	ID        int64     `json:"id"`
	UserName  string    `json:"username"`
	Password  string    `json:"password,omitempty"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type RegisterRequest struct {
    Username string `validate:"required,max=16"`
    Password string `validate:"required,min=8,max=16"`
    Name     string `validate:"required,min=2"`
}