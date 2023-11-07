package servers

import (
	"context"

	"github.com/gabrielmvnog/go-coupon/customer/src/models"
	"github.com/gabrielmvnog/go-coupon/customer/src/repositories"
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

func NewCustomerServiceServer(db *gorm.DB) *CustomerServiceServer {
	repository := repositories.NewCustomerRepository(db)
	usecase := usecases.NewCustomerUseCase(repository)

	return &CustomerServiceServer{usecase: usecase}
}

func (s *CustomerServiceServer) CreateCustomer(ctx context.Context, in *pb.CreateCustomerRequest) (*pb.CreateCustomerResponse, error) {
	customer := s.usecase.CreateCustomer(ctx, models.Customer{
		FirstName: in.FirstName,
		LastName:  in.LastName,
		Email:     in.Email,
		Phone:     in.Phone,
	})

	return &pb.CreateCustomerResponse{
		Id: customer.ID,
	}, nil
}

func (s *CustomerServiceServer) GetCustomer(ctx context.Context, in *pb.GetCustomerRequest) (*pb.GetCustomerResponse, error) {
	customer, err := s.usecase.GetCustomerById(ctx, in.Id)

	if err.Error() == "record not found" {
		return nil, status.Errorf(codes.NotFound, "user not found")
	}

	return &pb.GetCustomerResponse{
		FirstName: customer.FirstName,
		LastName:  customer.LastName,
		Email:     customer.Email,
		Phone:     customer.Phone,
	}, nil
}

func (s *CustomerServiceServer) UpdateCustomer(ctx context.Context, in *pb.UpdateCustomerRequest) (*pb.UpdateCustomerResponse, error) {
	customer := s.usecase.UpdateCustomer(ctx, models.Customer{
		ID:        in.Id,
		FirstName: in.FirstName,
		LastName:  in.LastName,
		Email:     in.Email,
		Phone:     in.Phone,
	})

	return &pb.UpdateCustomerResponse{
		Id:        customer.ID,
		FirstName: customer.FirstName,
		LastName:  customer.LastName,
		Email:     customer.Email,
		Phone:     customer.Phone,
	}, nil
}
