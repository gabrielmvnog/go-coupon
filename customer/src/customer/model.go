package customer

import "time"

type Customer struct {
	ID        uint
	FirstName string
	LastName  string
	Email     string
	Phone     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
