package usecases_test

import (
	"context"
	"testing"

	"github.com/gabrielmvnog/go-coupon/customer/src/models"
	"github.com/gabrielmvnog/go-coupon/customer/src/repositories/mocks"
	"github.com/gabrielmvnog/go-coupon/customer/src/usecases"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestCreateCustomer(t *testing.T) {
	type args struct {
		createCustomerObject models.Customer
		createCustomerResult models.Customer
		createCustomerError  error
	}
	customer := models.Customer{}
	tests := []struct {
		name      string
		args      args
		expect    models.Customer
		expectErr error
	}{
		{
			name: "should return success",
			args: args{
				createCustomerObject: models.Customer{},
				createCustomerResult: customer,
				createCustomerError:  nil,
			},
			expect:    customer,
			expectErr: nil,
		},
		{
			name: "should return error",
			args: args{
				createCustomerObject: models.Customer{},
				createCustomerResult: customer,
				createCustomerError:  gorm.ErrDuplicatedKey,
			},
			expectErr: gorm.ErrDuplicatedKey,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := new(mocks.CustomerRepository)
			r.On("CreateCustomer", context.Background(), &tt.args.createCustomerObject).Return(
				&tt.args.createCustomerResult, tt.args.createCustomerError,
			)

			u := usecases.NewCustomerUseCase(r)
			got, err := u.CreateCustomer(context.Background(), tt.args.createCustomerObject)

			if err != nil {
				assert.Equal(t, tt.expectErr, err)
			} else {
				assert.Equal(t, &tt.expect, got)
			}
		})

	}
}

func TestGetCustomerById(t *testing.T) {
	r := new(mocks.CustomerRepository)

	r.On("GetCustomerById", context.Background(), uint32(123)).Return(&models.Customer{}, nil)

	usecases.NewCustomerUseCase(r).GetCustomerById(context.Background(), 123)

	r.AssertExpectations(t)
}

func TestUpdateCustomer(t *testing.T) {
	r := new(mocks.CustomerRepository)
	customer := models.Customer{ID: uint32(123)}

	r.On("UpdateCustomer", context.Background(), uint32(123), &customer).Return(&customer, nil)

	usecases.NewCustomerUseCase(r).UpdateCustomer(context.Background(), customer)

	r.AssertExpectations(t)
}

func TestDeleteCustomer(t *testing.T) {
	r := new(mocks.CustomerRepository)

	r.On("DeleteCustomer", context.Background(), uint32(123)).Return(nil, nil)

	usecases.NewCustomerUseCase(r).DeleteCustomer(context.Background(), uint32(123))

	r.AssertExpectations(t)
}
