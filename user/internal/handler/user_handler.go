package handler

import (
	"context"
	"time"

	"user/proto/userpb"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type UserHandler struct {
	userClient userpb.UserServiceClient
}

func NewUserHandler(conn *grpc.ClientConn) *UserHandler {
	client := userpb.NewUserServiceClient(conn)
	return &UserHandler{
		userClient: client,
	}
}

// RegisterUser handles user registration
func (h *UserHandler) RegisterUser(c *gin.Context) {
	var req userpb.UserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "invalid request"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	resp, err := h.userClient.RegisterUser(ctx, &req)
	if err != nil {
		c.JSON(500, gin.H{"error": "gRPC error: " + err.Error()})
		return
	}

	c.JSON(200, resp)
}

// AuthenticateUser handles user authentication
func (h *UserHandler) AuthenticateUser(c *gin.Context) {
	var req userpb.AuthRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "invalid request"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	resp, err := h.userClient.AuthenticateUser(ctx, &req)
	if err != nil {
		c.JSON(500, gin.H{"error": "gRPC error: " + err.Error()})
		return
	}

	c.JSON(200, resp)
}

// GetUserProfile handles retrieving a user profile
func (h *UserHandler) GetUserProfile(c *gin.Context) {
	userID := c.Param("user_id")
	if userID == "" {
		c.JSON(400, gin.H{"error": "user_id is required"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	resp, err := h.userClient.GetUserProfile(ctx, &userpb.UserID{UserId: userID})
	if err != nil {
		c.JSON(500, gin.H{"error": "gRPC error: " + err.Error()})
		return
	}

	c.JSON(200, resp)
}
