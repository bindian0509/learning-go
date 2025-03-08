package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/bindian0509/learning-go/grpc/userpb"
	"google.golang.org/grpc"
)


func main() {
	// Connect to the gRPC server
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewUserServiceClient(conn) // ✅ Create gRPC client

	// Create a request
	req := &pb.UserRequest{Id: "123"}

	// Set timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Call gRPC method
	res, err := client.GetUser(ctx, req)
	if err != nil {
		log.Fatalf("Error calling GetUser: %v", err)
	}

	fmt.Printf("User Info: ID=%s, Name=%s, Email=%s\n", res.GetId(), res.GetName(), res.GetEmail())
}
