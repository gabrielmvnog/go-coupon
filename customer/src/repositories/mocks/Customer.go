package mocks

import (
	"context"

	"github.com/gabrielmvnog/go-coupon/customer/src/models"
	"github.com/stretchr/testify/mock"
)

type CustomerRepository struct {
	mock.Mock
}

func (_m *CustomerRepository) CreateCustomer(ctx context.Context, customer *models.Customer) (*models.Customer, error) {
	args := _m.Called(customer)

	return customer, args.Error(1)

}

func (_m *CustomerRepository) GetCustomerById(ctx context.Context, customer_id uint32) (*models.Customer, error) {
	var customer models.Customer
	args := _m.Called(customer)

	return &customer, args.Error(1)
}

func (_m *CustomerRepository) UpdateCustomer(ctx context.Context, customer_id uint32, customer *models.Customer) (*models.Customer, error) {
	args := _m.Called(customer)

	return customer, args.Error(1)
}
