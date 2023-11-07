package usecases_test

import (
	"context"
	"testing"

	"github.com/gabrielmvnog/go-coupon/customer/src/models"
	"github.com/gabrielmvnog/go-coupon/customer/src/repositories/mocks"
	"github.com/gabrielmvnog/go-coupon/customer/src/usecases"
	"github.com/stretchr/testify/mock"
)

func TestCreateCustomer(t *testing.T) {
	repository := new(mocks.CustomerRepository)

	repository.On("CreateCustomer", mock.Anything).Return(models.Customer{}, nil)

	usecases.NewCustomerUseCase(repository).CreateCustomer(context.Background(), models.Customer{})

	repository.AssertExpectations(t)
}

func TestGetCustomerById(t *testing.T) {
	repository := new(mocks.CustomerRepository)

	repository.On("GetCustomerById", mock.Anything).Return(models.Customer{}, nil)

	usecases.NewCustomerUseCase(repository).GetCustomerById(context.Background(), 123)

	repository.AssertExpectations(t)
}

func TestUpdateCustomer(t *testing.T) {
	repository := new(mocks.CustomerRepository)
	customer := models.Customer{ID: 1}

	repository.On("UpdateCustomer", mock.Anything).Return(customer, nil)

	usecases.NewCustomerUseCase(repository).UpdateCustomer(context.Background(), customer)

	repository.AssertExpectations(t)
}

func TestDeleteCustomer(t *testing.T) {
	repository := new(mocks.CustomerRepository)

	repository.On("DeleteCustomer", mock.Anything).Return(nil, nil)

	usecases.NewCustomerUseCase(repository).DeleteCustomer(context.Background(), 123)

	repository.AssertExpectations(t)
}
