package servers

import (
	"context"

	pb "github.com/gabrielmvnog/go-coupon/customer/proto"
	"github.com/gabrielmvnog/go-coupon/customer/src/models"
	"github.com/gabrielmvnog/go-coupon/customer/src/repositories"
	"github.com/gabrielmvnog/go-coupon/customer/src/usecases"
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
	customer := s.usecase.GetCustomerById(ctx, in.Id)

	return &pb.GetCustomerResponse{
		FirstName: customer.FirstName,
		LastName:  customer.LastName,
		Email:     customer.Email,
		Phone:     customer.Phone,
	}, nil
}
