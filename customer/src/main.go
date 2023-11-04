package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/gabrielmvnog/go-coupon/customer/proto"
	"google.golang.org/grpc"
)

var (
	host = flag.String("host", "0.0.0.0", "Server host")
	port = flag.Int("port", 9000, "Server port")
)

type customerServer struct {
	pb.UnimplementedCustomerServiceServer
}

func (s *customerServer) CreateCustomer(ctx context.Context, in *pb.Customer) (*pb.CustomerIdentifier, error) {
	log.Printf("Receive: %v", in.GetPhone())
	return &pb.CustomerIdentifier{}, nil
}

func (s *customerServer) GetCustomer(ctx context.Context, in *pb.CustomerIdentifier) (*pb.CustomerResponse, error) {
	log.Printf("Receive: %v", in.GetId())
	return &pb.CustomerResponse{}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *host, *port))
	if err != nil {
		log.Fatalf("failad to listen: %v", err)
	}

	server := grpc.NewServer()
	pb.RegisterCustomerServiceServer(server, &customerServer{})

	log.Printf("server listening at %v", lis.Addr())
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
