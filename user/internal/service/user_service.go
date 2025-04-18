package service

import (
	"context"
	"fmt"
	"user/proto/userpb"
)

type UserServer struct {
	userpb.UnimplementedUserServiceServer
}

func (s *UserServer) RegisterUser(ctx context.Context, req *userpb.UserRequest) (*userpb.UserResponse, error) {
	fmt.Println("RegisterUser called with:", req.Username)
	return &userpb.UserResponse{
		Message: "User registered successfully",
		UserId:  "12345",
	}, nil
}

func (s *UserServer) AuthenticateUser(ctx context.Context, req *userpb.AuthRequest) (*userpb.AuthResponse, error) {
	fmt.Println("AuthenticateUser called with:", req.Username)
	return &userpb.AuthResponse{
		Success: true,
		Token:   "token123",
	}, nil
}

func (s *UserServer) GetUserProfile(ctx context.Context, req *userpb.UserID) (*userpb.UserProfile, error) {
	fmt.Println("GetUserProfile called with:", req.UserId)
	return &userpb.UserProfile{
		UserId:   req.UserId,
		Username: "demo_user",
	}, nil
}
