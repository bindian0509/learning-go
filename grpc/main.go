package main

import (
	"fmt"
	"log"
	"net"

	"github.com/bindian0509/learning-go/grpc/server"
	pb "github.com/bindian0509/learning-go/grpc/userpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"  // ✅ Import reflection

)

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, &server.UserServiceServer{})

	// ✅ Enable Reflection
	reflection.Register(grpcServer)

	fmt.Println("gRPC Server is running on port 50051...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
