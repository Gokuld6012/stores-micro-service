package main

import (
	"context"
	"log"
	"net"

	pb "github.com/gokuld6012/product-micro/product-service/proto"
	"google.golang.org/grpc"
)

const (
	port = ":50052"
)

type service struct{}

func (s *service) GetTestProductResponse(ctx context.Context, c *pb.ProductTestRequest) (*pb.ProductTestResponse, error) {
	return &pb.ProductTestResponse{
		Message: "This is from product service",
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterProductServer(s, &service{})

	log.Println("Running on port", port)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
