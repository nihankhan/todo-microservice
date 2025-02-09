package handlers

import (
	"context"
	"todo/internal/service"
	"todo/proto"

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

	completed := todo.Completed

	return &proto.TodoResponse{
		Todo: &proto.Todo{
			Id:          todo.ID,
			Title:       todo.Title,
			Description: todo.Description,
			Completed:   &completed,
		},
	}, nil
}

func (h *GrpcHandler) CreateTodo(ctx context.Context, r *proto.CreateTodoRequest) (*proto.CreateTodoResponse, error) {
	completed := false
	todo, err := h.s.CreatedTodo(r.Title, r.Description, completed)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error create todo: %v", err)
	}

	return &proto.CreateTodoResponse{
		Todo: &proto.Todo{
			Id:          todo.ID,
			Title:       todo.Title,
			Description: todo.Description,
		},
	}, nil
}

func (h *GrpcHandler) UpdateTodo(ctx context.Context, r *proto.UpdateTodoRequest) (*proto.UpdateTodoResponse, error) {
	completed := false

	todo, err := h.s.UpdateTodo(r.TodoId, r.Title, r.Description, completed)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error create todo: %v", err)
	}

	return &proto.UpdateTodoResponse{
		Todo: &proto.Todo{
			Id:          todo.ID,
			Title:       todo.Title,
			Description: todo.Description,
			Completed:   &completed,
		},
	}, nil
}

func (h *GrpcHandler) DeleteTodo(ctx context.Context, r *proto.DeleteTodoRequest) (*proto.DeleteTodoResponse, error) {
	err := h.s.DeleteTodo(r.TodoId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error delete todo: %v", err)
	}

	// Create a bool pointer
	status := true

	return &proto.DeleteTodoResponse{
		Status:  &status,
		Message: "Todo deleted successfully",
	}, nil
}

func (h *GrpcHandler) GetAllTodos(ctx context.Context, r *proto.GetAllTodosRequest) (*proto.GetAllTodosResponse, error) {
	todos, err := h.s.GetAllTodos()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error getting all todos: %v", err)
	}

	var todosResponse []*proto.Todo
	for _, todo := range todos {
		todosResponse = append(todosResponse, &proto.Todo{
			Id:          todo.ID,
			Title:       todo.Title,
			Description: todo.Description,
			Completed:   boolPtr(todo.Completed),
		})
	}

	// Return the todos as a response
	return &proto.GetAllTodosResponse{
		Todo: todosResponse,
	}, nil
}

func (h *GrpcHandler) MarkAsDone(ctx context.Context, r *proto.MarkRequest) (*proto.MarkResponse, error) {
	_, err := h.s.MarkAsDone(r.TodoId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error marking todo as completed: %v", err)
	}

	return &proto.MarkResponse{
		Message: "Marked as Done.",
	}, nil
}

func boolPtr(b bool) *bool {
	return &b
}
