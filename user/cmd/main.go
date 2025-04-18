package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"log"
	"net"
	"sync"
	"user/internal/handler"
	"user/internal/service"
	"user/proto/userpb"
)

func startGRPCServer() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	userpb.RegisterUserServiceServer(s, &service.UserServer{})

	fmt.Println("User service is running on port :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func startHTTPServer() {
	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	userHandler := handler.NewUserHandler(conn)

	r := gin.Default()

	// Define routes
	r.POST("/users/register", userHandler.RegisterUser)
	r.POST("/users/authenticate", userHandler.AuthenticateUser)
	r.GET("/users/:user_id", userHandler.GetUserProfile)

	// Start HTTP server on port 8083
	r.Run(":8083")
}

func main() {
	// Create a WaitGroup to wait for both servers to finish
	var wg sync.WaitGroup
	wg.Add(2)

	// Start gRPC server in a goroutine
	go func() {
		defer wg.Done()
		startGRPCServer()
	}()

	// Start HTTP server in a goroutine
	go func() {
		defer wg.Done()
		startHTTPServer()
	}()

	// Wait for both servers to finish
	wg.Wait()
}
