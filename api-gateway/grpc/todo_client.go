package grpc

import (
	"api-gateway/proto"
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
)

type TodoClient struct {
	client proto.TodoServiceClient
	conn   *grpc.ClientConn
}

func NewTodoClient() *TodoClient {
	conn, err := grpc.Dial("127.0.0.1:50053", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect to Todo service: %v", err)
	}

	client := proto.NewTodoServiceClient(conn)

	fmt.Println("client: ", client)

	return &TodoClient{
		client: client,
		conn:   conn,
	}

}

func (c *TodoClient) GetTodoByID(ctx context.Context, todoID string) (*proto.TodoResponse, error) {
	r := &proto.TodoRequest{
		TodoId: todoID,
	}

	return c.client.GetTodo(ctx, r)
}

func (c *TodoClient) CreateTodo(ctx context.Context, todo *proto.CreateTodoRequest) (*proto.CreateTodoResponse, error) {
	resp, err := c.client.CreateTodo(ctx, todo)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *TodoClient) UpdateTodo(ctx context.Context, todo *proto.UpdateTodoRequest) (*proto.UpdateTodoResponse, error) {
	resp, err := c.client.UpdateTodo(ctx, todo)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *TodoClient) DeleteTodo(ctx context.Context, todoID string) (*proto.DeleteTodoResponse, error) {
	r := &proto.DeleteTodoRequest{
		TodoId: todoID,
	}

	return c.client.DeleteTodo(ctx, r)
}
