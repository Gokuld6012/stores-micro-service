package main

import (
	"log"
	"net"
	"context"
	pb "github.com/gokuld6012/product-micro/auth-service/proto"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type service struct {}

func (s *service) GetTestResponse(ctx context.Context, c *pb.TestRequest) (*pb.TestResponse, error) {
	return &pb.TestResponse{
		Message: "This is from auth service",
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterAuthServer(s, &service{})

	log.Println("Running on port", port)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
