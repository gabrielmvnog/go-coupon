package customer

import (
	"context"
	"log"

	"gorm.io/gorm"
)

type CustomerUseCase struct {
	repository CustomerRepository
}

func NewCustomerUseCase(db *gorm.DB) *CustomerUseCase {
	repository := NewCustomerRepository(db)

	return &CustomerUseCase{
		repository: *repository,
	}
}

func (u *CustomerUseCase) GetCustomerById(ctx context.Context, customer_id uint) *Customer {
	log.Printf("Getting user: %d", customer_id)
	customer, _ := u.repository.GetCustomerById(ctx, customer_id)

	return &customer
}
