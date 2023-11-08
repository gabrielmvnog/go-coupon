package usecases

import (
	"context"
	"log"

	"github.com/gabrielmvnog/go-coupon/customer/src/models"
	"github.com/gabrielmvnog/go-coupon/customer/src/repositories"
	"gorm.io/gorm"
)

//go:generate mockery --name CustomerUseCase
type CustomerUseCase interface {
	CreateCustomer(ctx context.Context, customer models.Customer) (*models.Customer, error)
	GetCustomerById(ctx context.Context, customer_id uint32) (*models.Customer, error)
	UpdateCustomer(ctx context.Context, customer models.Customer) *models.Customer
	DeleteCustomer(ctx context.Context, customer_id uint32) error
}

type customerUseCase struct {
	repository repositories.CustomerRepository
}

func NewCustomerUseCase(repository repositories.CustomerRepository) *customerUseCase {
	return &customerUseCase{
		repository: repository,
	}
}

func (u *customerUseCase) CreateCustomer(ctx context.Context, customer models.Customer) (*models.Customer, error) {
	log.Printf("Creating user: %s", customer.FirstName)

	u.repository.CreateCustomer(ctx, &customer)

	return &customer, nil
}

func (u *customerUseCase) GetCustomerById(ctx context.Context, customer_id uint32) (*models.Customer, error) {
	log.Printf("Getting user: %d", customer_id)
	customer, err := u.repository.GetCustomerById(ctx, customer_id)

	if err == gorm.ErrRecordNotFound {
		return nil, err
	}

	return customer, nil
}

func (u *customerUseCase) UpdateCustomer(ctx context.Context, customer models.Customer) *models.Customer {
	log.Printf("Updating user: %d", customer.ID)

	u.repository.UpdateCustomer(ctx, customer.ID, &customer)

	return &customer
}

func (u *customerUseCase) DeleteCustomer(ctx context.Context, customer_id uint32) error {
	log.Printf("Deleting user: %d", customer_id)

	err := u.repository.DeleteCustomer(ctx, customer_id)

	return err
}
