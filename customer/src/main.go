package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/gabrielmvnog/go-coupon/customer/proto"
	db "github.com/gabrielmvnog/go-coupon/customer/src/db"
	"github.com/gabrielmvnog/go-coupon/customer/src/servers"
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

	customerServiceServer := servers.NewCustomerServiceServer(db)
	pb.RegisterCustomerServiceServer(server, customerServiceServer)

	log.Printf("server listening at %v", lis.Addr())
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
