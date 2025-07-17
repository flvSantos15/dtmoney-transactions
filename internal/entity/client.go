package entity

import "time"

type Client struct {
	ID string
	Name string
	Email string
	CreatedAt time.Time
	UpdatedAt time.Time
}

