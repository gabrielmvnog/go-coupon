package servers

import (
	"context"

	"github.com/gabrielmvnog/go-coupon/customer/src/models"
	pb "github.com/gabrielmvnog/go-coupon/customer/src/servers/proto"
	"github.com/gabrielmvnog/go-coupon/customer/src/usecases"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"gorm.io/gorm"
)

type CustomerServiceServer struct {
	pb.UnimplementedCustomerServiceServer

	usecase usecases.CustomerUseCase
}

func NewCustomerServiceServer(usecase usecases.CustomerUseCase) *CustomerServiceServer {
	return &CustomerServiceServer{usecase: usecase}
}

func (s *CustomerServiceServer) CreateCustomer(ctx context.Context, in *pb.CreateCustomerRequest) (*pb.CreateCustomerResponse, error) {
	body := models.Customer{
		FirstName: in.FirstName,
		LastName:  in.LastName,
		Email:     in.Email,
		Phone:     in.Phone,
	}
	customer := s.usecase.CreateCustomer(ctx, &body)

	return &pb.CreateCustomerResponse{
		Id: customer.ID,
	}, nil
}

func (s *CustomerServiceServer) GetCustomer(ctx context.Context, in *pb.GetCustomerRequest) (*pb.GetCustomerResponse, error) {
	customer, err := s.usecase.GetCustomerById(ctx, in.Id)

	if err == gorm.ErrRecordNotFound {
		return nil, status.Errorf(codes.NotFound, "user not found")
	}

	return &pb.GetCustomerResponse{
		FirstName: customer.FirstName,
		LastName:  customer.LastName,
		Email:     customer.Email,
		Phone:     customer.Phone,
	}, nil
}

func (s *CustomerServiceServer) DeleteCustomer(ctx context.Context, in *pb.DeleteCustomerRequest) (*pb.DeleteCustomerResponse, error) {
	s.usecase.DeleteCustomer(ctx, in.Id)

	return &pb.DeleteCustomerResponse{}, nil
}

func (s *CustomerServiceServer) UpdateCustomer(ctx context.Context, in *pb.UpdateCustomerRequest) (*pb.UpdateCustomerResponse, error) {
	body := models.Customer{
		ID:        in.Id,
		FirstName: in.FirstName,
		LastName:  in.LastName,
		Email:     in.Email,
		Phone:     in.Phone,
	}
	customer := s.usecase.UpdateCustomer(ctx, &body)

	return &pb.UpdateCustomerResponse{
		Id:        customer.ID,
		FirstName: customer.FirstName,
		LastName:  customer.LastName,
		Email:     customer.Email,
		Phone:     customer.Phone,
	}, nil
}
