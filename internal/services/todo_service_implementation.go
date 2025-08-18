package services

import (
	"context"
	"todo-list-be/internal/domain"
	"todo-list-be/internal/dto"
	"todo-list-be/internal/repository"
)

type TodoServiceImplementation struct{
	repo repository.TodoRepository
}

func NewTodoService(repo repository.TodoRepository) TodoService{
	return &TodoServiceImplementation{repo: repo}
}

func (s *TodoServiceImplementation) CreateTodo(request *dto.CreateTodoRequest) (domain.Todo, error){
	if request.Priority == "" {
		request.Priority = "Medium"
	}

	return s.repo.Save(request)
}

func (s *TodoServiceImplementation) UpdateTodo(ctx context.Context, request *dto.UpdateTodoRequest) (domain.Todo, error){
	return s.repo.Update(ctx, request)
}