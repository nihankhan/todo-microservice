package handlers

import (
	"api-gateway/proto"
	"context"
	"todo/internal/service"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GrpcHandler struct {
	proto.UnimplementedTodoServiceServer
	s *service.TodoService
}

func NewGrpcHandler(service *service.TodoService) *GrpcHandler {
	return &GrpcHandler{
		s: service,
	}
}

func (h *GrpcHandler) GetTodo(ctx context.Context, r *proto.TodoRequest) (*proto.TodoResponse, error) {
	todo, err := h.s.GetTodoByID(r.TodoId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error getting todo: %v", err)
	}

	return &proto.TodoResponse{
		Todo: &proto.Todo{
			Id:          todo.ID,
			Title:       todo.Title,
			Description: todo.Description,
			Completed:   todo.Completed,
		},
	}, nil
}

func (h *GrpcHandler) CreateTodo(ctx context.Context, r *proto.CreateTodoRequest) (*proto.CreateTodoResponse, error) {
	todo, err := h.s.CreatedTodo(r.Title, r.Description)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error create todo: %v", err)
	}

	return &proto.CreateTodoResponse{
		Todo: &proto.Todo{
			Id:          todo.ID,
			Title:       todo.Title,
			Description: todo.Description,
			Completed:   todo.Completed,
		},
	}, nil
}

func (h GrpcHandler) UpdateTodo(ctx context.Context, r *proto.UpdateTodoRequest) (*proto.UpdateTodoResponse, error) {
	todo, err := h.s.UpdateTodo(r.TodoId, r.Title, r.Description, r.Completed)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error create todo: %v", err)
	}

	return &proto.UpdateTodoResponse{
		Todo: &proto.Todo{
			Id:          todo.ID,
			Title:       todo.Title,
			Description: todo.Description,
			Completed:   todo.Completed,
		},
	}, nil
}

func (h *GrpcHandler) DeleteTodo(ctx context.Context, r *proto.DeleteTodoRequest) (*proto.DeleteTodoResponse, error) {
	err := h.s.DeleteTodo(r.TodoId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error delete todo: %v", err)
	}

	return &proto.DeleteTodoResponse{
		Status:  true,
		Message: "Todo deleted successfully",
	}, nil
}
