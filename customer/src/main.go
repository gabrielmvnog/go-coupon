package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/gabrielmvnog/go-coupon/customer/proto"
	customer "github.com/gabrielmvnog/go-coupon/customer/src/customer"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host = flag.String("host", "0.0.0.0", "Server host")
	port = flag.Int("port", 9000, "Server port")
)

func main() {
	flag.Parse()

	db := initDb()

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *host, *port))
	if err != nil {
		log.Fatalf("failad to listen: %v", err)
	}

	server := grpc.NewServer()

	customerServiceServer := customer.NewCustomerServiceServer(db)
	pb.RegisterCustomerServiceServer(server, customerServiceServer)

	log.Printf("server listening at %v", lis.Addr())
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func initDb() *gorm.DB {
	dsn := "postgres://customer:customer@0.0.0.0:5432/customer"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	return db
}
