package grpc

import (
	"api-gateway/proto"
	"context"
	"fmt"

	"google.golang.org/grpc"
)

type UserServiceClient struct {
	client proto.UserServiceClient
	conn   *grpc.ClientConn
}

func NewUserServiceClient() *UserServiceClient {
	conn, err := grpc.Dial(":50052", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("could not connect to User Service: %v", err)
	}

	client := proto.NewUserServiceClient(conn)

	return &UserServiceClient{
		client: client,
		conn:   conn,
	}
}

func (c *UserServiceClient) GetUserByID(ctx context.Context, userID string) (*proto.UserResponse, error) {
	r := &proto.UserRequest{
		UserId: userID,
	}

	return c.client.GetUser(ctx, r)
}

func (c *UserServiceClient) Register(ctx context.Context, user *proto.RegisterRequest) (*proto.RegisterResponse, error) {
	resp, err := c.client.Register(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("could not create user: %w", err)
	}

	return resp, nil
}

func (c *UserServiceClient) Login(ctx context.Context, login *proto.LoginRequest) (*proto.LoginResponse, error) {
	resp, err := c.client.Login(ctx, login)
	if err != nil {
		return nil, fmt.Errorf("could not logging in: %w", err)
	}

	return resp, nil
}

func (c *UserServiceClient) Logout(ctx context.Context, logout *proto.LogoutRequest) (*proto.LogoutResponse, error) {
	resp, err := c.client.Logout(ctx, logout)
	if err != nil {
		return nil, fmt.Errorf("could not logged out: %w", err)
	}

	return resp, nil
}

func (c *UserServiceClient) UpdateUser(ctx context.Context, user *proto.UpdateUserRequest) (*proto.UpdateUserResponse, error) {
	resp, err := c.client.UpdateUser(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("could not update user: %w", err)
	}

	return resp, nil

}

func (c *UserServiceClient) DeleteUser(ctx context.Context, userID string) (*proto.DeleteUserResponse, error) {
	r := &proto.DeleteUserRequest{
		UserId: userID,
	}

	return c.client.DeleteUser(ctx, r)
}
