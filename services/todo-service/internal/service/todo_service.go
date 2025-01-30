package service

import (
	"fmt"
	"time"
	"todo/internal/domain"
	"todo/internal/repository"
)

type TodoService struct {
	repo repository.TodoRepository
}

func NewTodoService(repo repository.TodoRepository) *TodoService {
	return &TodoService{
		repo: repo,
	}
}

func (s *TodoService) GetTodoByID(ID string) (*domain.Todo, error) {
	todo, err := s.repo.GetTodoByID(ID)
	if err != nil {
		return nil, fmt.Errorf("could not fetch todo: %w", err)
	}

	return todo, nil
}

func (s *TodoService) CreatedTodo(title, description string) (*domain.Todo, error) {
	todo := &domain.Todo{
		ID:          generateID(),
		Title:       title,
		Description: description,
	}

	todo, err := s.repo.CreatedTodo(todo)
	if err != nil {
		return nil, fmt.Errorf("could not create user: %w", err)
	}

	return todo, nil
}

func (s *TodoService) UpdateTodo(ID, title, description string, completed bool) (*domain.Todo, error) {
	todo := &domain.Todo{
		ID:          ID,
		Title:       title,
		Description: description,
		Completed:   completed,
	}

	todo, err := s.repo.UpdateTodo(todo)
	if err != nil {
		return nil, fmt.Errorf("failed to update todo: %w", err)
	}

	return todo, nil
}

func (s *TodoService) DeleteTodo(ID string) error {
	err := s.repo.DeleteTodo(ID)
	if err != nil {
		return fmt.Errorf("failed to delete todo: %w", err)
	}

	return nil
}

func generateID() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}
