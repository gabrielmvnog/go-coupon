package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	db "github.com/gabrielmvnog/go-coupon/customer/src/db"
	"github.com/gabrielmvnog/go-coupon/customer/src/repositories"
	"github.com/gabrielmvnog/go-coupon/customer/src/servers"
	pb "github.com/gabrielmvnog/go-coupon/customer/src/servers/proto"
	"github.com/gabrielmvnog/go-coupon/customer/src/usecases"
	"google.golang.org/grpc"
)

var (
	host = flag.String("host", "0.0.0.0", "Server host")
	port = flag.Int("port", 9000, "Server port")
)

func main() {
	flag.Parse()

	db := db.InitDB()

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *host, *port))
	if err != nil {
		log.Fatalf("failad to listen: %v", err)
	}

	server := grpc.NewServer()

	customerRepository := repositories.NewCustomerRepository(db)
	customerUsecase := usecases.NewCustomerUseCase(customerRepository)
	customerServiceServer := servers.NewCustomerServiceServer(customerUsecase)
	pb.RegisterCustomerServiceServer(server, customerServiceServer)

	log.Printf("server listening at %v", lis.Addr())
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
