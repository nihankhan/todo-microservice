package repository

import (
	"fmt"
	"todo/internal/domain"
)

type TodoRepository interface {
	GetTodoByID(ID string) (*domain.Todo, error)
	CreatedTodo(todo *domain.Todo) (*domain.Todo, error)
	UpdateTodo(todo *domain.Todo) (*domain.Todo, error)
	DeleteTodo(ID string) error
	GetAllTodos() ([]domain.Todo, error)
	MarkAsDone(ID string) (bool, error)
}

type InMTodoRepo struct {
	todo map[string]*domain.Todo
}

func NewTodoRepository() *InMTodoRepo {
	return &InMTodoRepo{
		todo: make(map[string]*domain.Todo),
	}
}

func (repo *InMTodoRepo) GetTodoByID(ID string) (*domain.Todo, error) {
	todo, exists := repo.todo[ID]
	if !exists {
		return nil, fmt.Errorf("todo not found")
	}

	return todo, nil
}

func (repo *InMTodoRepo) CreatedTodo(todo *domain.Todo) (*domain.Todo, error) {
	repo.todo[todo.ID] = todo

	return todo, nil
}

func (repo *InMTodoRepo) UpdateTodo(todo *domain.Todo) (*domain.Todo, error) {
	repo.todo[todo.ID] = todo

	return todo, nil
}

func (repo *InMTodoRepo) DeleteTodo(ID string) error {
	delete(repo.todo, ID)

	return nil
}

func (repo *InMTodoRepo) GetAllTodos() ([]domain.Todo, error) {
	todos := make([]domain.Todo, 0, len(repo.todo))

	for _, todo := range repo.todo {
		todos = append(todos, *todo)
	}

	return todos, nil
}

func (repo *InMTodoRepo) MarkAsDone(ID string) (bool, error) {
	todo, err := repo.GetTodoByID(ID)
	if err != nil {
		return false, err
	}

	todo.Completed = true

	return true, nil
}
