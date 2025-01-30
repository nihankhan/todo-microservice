package handlers

import (
	"context"
	"fmt"

	"user-service/internal/service"
	"user-service/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GrpcHandler struct {
	proto.UnimplementedUserServiceServer
	s *service.UserService
}

func NewGrpcHandler(service *service.UserService) *GrpcHandler {
	return &GrpcHandler{
		s: service,
	}
}

func (h *GrpcHandler) Register(ctx context.Context, r *proto.RegisterRequest) (*proto.RegisterResponse, error) {
	fmt.Println("create request")

	user, err := h.s.CreateUser(r.Name, r.Username, r.Email, r.Password)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error creating user: %v", err)
	}

	return &proto.RegisterResponse{
		User: &proto.User{
			Id:       user.ID,
			Name:     user.Name,
			Username: user.Username,
			Email:    user.Email,
			Password: user.Password,
			// ProfileImage: user.ProfileImage,
		},
	}, nil
}

func (h *GrpcHandler) Login(ctx context.Context, r *proto.LoginRequest) (*proto.LoginResponse, error) {
	token, err := h.s.Login(r.Email, r.Password)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Error logging in: %v", err)
	}

	return &proto.LoginResponse{
		Token: token,
	}, nil
}

func (h *GrpcHandler) GetUser(ctx context.Context, r *proto.UserRequest) (*proto.UserResponse, error) {
	fmt.Println("get request")

	user, err := h.s.GetUser(r.UserId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error getting user: %v", err)
	}

	return &proto.UserResponse{
		User: &proto.User{
			Id:           user.ID,
			Name:         user.Name,
			Username:     user.Username,
			Email:        user.Email,
			Password:     user.Password,
			ProfileImage: user.ProfileImage,
		},
	}, nil
}

func (h *GrpcHandler) UpdateUser(ctx context.Context, r *proto.UpdateUserRequest) (*proto.UpdateUserResponse, error) {
	fmt.Println("update request")

	user, err := h.s.UpdateUser(r.UserId, r.Name, r.Username, r.Email, r.Password, r.ProfileImage)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error creating user: %v", err)
	}

	return &proto.UpdateUserResponse{
		User: &proto.User{
			Id:           user.ID,
			Name:         user.Name,
			Username:     user.Username,
			Email:        user.Email,
			Password:     user.Password,
			ProfileImage: user.ProfileImage,
		},
	}, nil
}

func (h *GrpcHandler) DeleteUser(ctx context.Context, r *proto.DeleteUserRequest) (*proto.DeleteUserResponse, error) {
	fmt.Println("delete request")

	err := h.s.DeleteUser(r.UserId)
	if err != nil {
		return nil, status.Errorf(codes.Code(13), "Error deleting user: %v", err)
	}

	return &proto.DeleteUserResponse{
		Success: true,
		Message: "User deleted successfully",
	}, nil
}
