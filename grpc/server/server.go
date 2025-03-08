package server

import (
	"context"
	"log"
	pb "github.com/bindian0509/learning-go/grpc/userpb"
)

// Implement the gRPC service
type UserServiceServer struct {
	pb.UnimplementedUserServiceServer
}

func (s *UserServiceServer) GetUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	log.Printf("Received request for user ID: %s", req.GetId())

	// Mock response
	return &pb.UserResponse{
		Id:    req.GetId(),
		Name:  "John Doe",
		Email: "john@example.com",
	}, nil
}
