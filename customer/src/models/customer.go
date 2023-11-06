package models

import "time"

type Customer struct {
	ID        uint32
	FirstName string
	LastName  string
	Email     string
	Phone     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
