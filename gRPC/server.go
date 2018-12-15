package main

import (
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "github.com/shijuvar/go-recipes/grpc/customer"
)

// CreateCustomer creates a new Customer
func (s *server) getMessage(ctx context.Context, in *pb.testRequest) (*pb.testResponse, error) {
	return &pb.testResponse{text: "Hello world from gRPC"}, nil
}

func main() {
	lis, err := net.Listen("tcp", 3000)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// Creates a new gRPC server
	s := grpc.NewServer()
	pb.RegisterCustomerServer(s, &server{})
	s.Serve(lis)
}
