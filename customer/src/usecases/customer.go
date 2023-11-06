package usecases

import (
	"context"
	"log"

	"github.com/gabrielmvnog/go-coupon/customer/src/models"
	"github.com/gabrielmvnog/go-coupon/customer/src/repositories"
	"gorm.io/gorm"
)

type CustomerUseCase struct {
	repository repositories.CustomerRepository
}

func NewCustomerUseCase(db *gorm.DB) *CustomerUseCase {
	repository := repositories.NewCustomerRepository(db)

	return &CustomerUseCase{
		repository: *repository,
	}
}

func (u *CustomerUseCase) CreateCustomer(ctx context.Context, customer models.Customer) *models.Customer {
	log.Printf("Creating user: %s", customer.FirstName)

	u.repository.CreateCustomer(ctx, &customer)

	return &customer
}

func (u *CustomerUseCase) GetCustomerById(ctx context.Context, customer_id uint32) *models.Customer {
	log.Printf("Getting user: %d", customer_id)
	customer, _ := u.repository.GetCustomerById(ctx, customer_id)

	return customer
}
