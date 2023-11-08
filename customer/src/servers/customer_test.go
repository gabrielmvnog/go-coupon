package servers_test

import (
	"context"
	"testing"

	"github.com/gabrielmvnog/go-coupon/customer/src/models"
	"github.com/gabrielmvnog/go-coupon/customer/src/servers"
	pb "github.com/gabrielmvnog/go-coupon/customer/src/servers/proto"
	"github.com/gabrielmvnog/go-coupon/customer/src/usecases/mocks"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

func TestCreateCustomer(t *testing.T) {
	customer := models.Customer{}

	body := pb.CreateCustomerRequest{}
	expect := pb.CreateCustomerResponse{}

	u := new(mocks.CustomerUseCase)
	u.On("CreateCustomer", context.Background(), &customer).Return(
		&customer,
	)

	s := servers.NewCustomerServiceServer(u)
	got, _ := s.CreateCustomer(context.Background(), &body)

	assert.Equal(t, &expect, got)
}

func TestGetCustomer(t *testing.T) {
	customer := models.Customer{}

	type args struct {
		getCustomerResult models.Customer
		getCustomerError  error
	}
	tests := []struct {
		name      string
		args      args
		expect    pb.GetCustomerResponse
		expectErr error
	}{
		{
			name: "should return success",
			args: args{
				getCustomerResult: customer,
			},
			expect: pb.GetCustomerResponse{},
		},
		{
			name: "should return not found error",
			args: args{
				getCustomerResult: customer,
				getCustomerError:  gorm.ErrRecordNotFound,
			},
			expectErr: status.Errorf(codes.NotFound, "user not found"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body := pb.GetCustomerRequest{Id: uint32(123)}

			u := new(mocks.CustomerUseCase)
			u.On("GetCustomerById", context.Background(), uint32(123)).Return(
				&tt.args.getCustomerResult, tt.args.getCustomerError,
			)

			s := servers.NewCustomerServiceServer(u)
			got, err := s.GetCustomer(context.Background(), &body)

			if err != nil {
				assert.Equal(t, tt.expectErr, err)
			} else {
				assert.Equal(t, &tt.expect, got)
			}
		})
	}
}

func TestUpdateCustomer(t *testing.T) {
	customer := models.Customer{}

	body := pb.UpdateCustomerRequest{}
	expect := pb.UpdateCustomerResponse{}

	u := new(mocks.CustomerUseCase)
	u.On("UpdateCustomer", context.Background(), &customer).Return(
		&customer,
	)

	s := servers.NewCustomerServiceServer(u)
	got, _ := s.UpdateCustomer(context.Background(), &body)

	assert.Equal(t, &expect, got)
}

func TestDeleteCustomer(t *testing.T) {
	body := pb.DeleteCustomerRequest{
		Id: uint32(123),
	}
	expect := pb.DeleteCustomerResponse{}

	u := new(mocks.CustomerUseCase)
	u.On("DeleteCustomer", context.Background(), uint32(123)).Return(
		nil,
	)

	s := servers.NewCustomerServiceServer(u)
	got, _ := s.DeleteCustomer(context.Background(), &body)

	assert.Equal(t, &expect, got)
}
