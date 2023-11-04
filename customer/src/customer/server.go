package customer

import (
	"context"

	pb "github.com/gabrielmvnog/go-coupon/customer/proto"
	"gorm.io/gorm"
)

type CustomerServiceServer struct {
	pb.UnimplementedCustomerServiceServer

	usecase CustomerUseCase
}

func NewCustomerServiceServer(db *gorm.DB) *CustomerServiceServer {
	usecase := NewCustomerUseCase(db)

	return &CustomerServiceServer{usecase: *usecase}
}

func (s *CustomerServiceServer) CreateCustomer(ctx context.Context, in *pb.Customer) (*pb.CustomerIdentifier, error) {
	return &pb.CustomerIdentifier{}, nil
}

func (s *CustomerServiceServer) GetCustomer(ctx context.Context, in *pb.CustomerIdentifier) (*pb.CustomerResponse, error) {
	customer := s.usecase.GetCustomerById(ctx, 1)

	return &pb.CustomerResponse{
		// id:        customer.ID,
		FirstName: customer.FirstName,
		LastName:  customer.LastName,
		Email:     customer.Email,
		Phone:     customer.Phone,
	}, nil
}
