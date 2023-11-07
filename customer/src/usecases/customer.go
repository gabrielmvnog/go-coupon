package usecases

import (
	"context"
	"log"

	"github.com/gabrielmvnog/go-coupon/customer/src/models"
	"github.com/gabrielmvnog/go-coupon/customer/src/repositories"
)

type CustomerUseCase interface {
	CreateCustomer(ctx context.Context, customer models.Customer) *models.Customer
	GetCustomerById(ctx context.Context, customer_id uint32) *models.Customer
	UpdateCustomer(ctx context.Context, customer models.Customer) *models.Customer
}

type customerUseCase struct {
	repository repositories.CustomerRepository
}

func NewCustomerUseCase(repository repositories.CustomerRepository) *customerUseCase {
	return &customerUseCase{
		repository: repository,
	}
}

func (u *customerUseCase) CreateCustomer(ctx context.Context, customer models.Customer) *models.Customer {
	log.Printf("Creating user: %s", customer.FirstName)

	u.repository.CreateCustomer(ctx, &customer)

	return &customer
}

func (u *customerUseCase) GetCustomerById(ctx context.Context, customer_id uint32) *models.Customer {
	log.Printf("Getting user: %d", customer_id)
	customer, _ := u.repository.GetCustomerById(ctx, customer_id)

	return customer
}

func (u *customerUseCase) UpdateCustomer(ctx context.Context, customer models.Customer) *models.Customer {
	log.Printf("Updating user: %d", customer.ID)

	u.repository.UpdateCustomer(ctx, customer.ID, &customer)

	return &customer
}
