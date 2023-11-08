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
	customer := models.Customer{}

	r := new(mocks.CustomerRepository)
	r.On("CreateCustomer", context.Background(), &customer).Return(
		&customer, nil,
	)

	u := usecases.NewCustomerUseCase(r)
	got, _ := u.CreateCustomer(context.Background(), customer)

	assert.Equal(t, &customer, got)
}

func TestGetCustomerById(t *testing.T) {
	type args struct {
		getCustomerResult models.Customer
		getCustomerError  error
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
				getCustomerResult: customer,
				getCustomerError:  nil,
			},
			expect:    customer,
			expectErr: nil,
		},
		{
			name: "should return error",
			args: args{
				getCustomerResult: customer,
				getCustomerError:  gorm.ErrRecordNotFound,
			},
			expectErr: gorm.ErrRecordNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			iD := uint32(123)
			r := new(mocks.CustomerRepository)

			r.On("GetCustomerById", context.Background(), iD).Return(&tt.args.getCustomerResult, tt.args.getCustomerError)

			u := usecases.NewCustomerUseCase(r)
			got, err := u.GetCustomerById(context.Background(), iD)

			if err != nil {
				assert.Equal(t, tt.expectErr, err)
			} else {
				assert.Equal(t, &tt.expect, got)
			}
		})

	}
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
