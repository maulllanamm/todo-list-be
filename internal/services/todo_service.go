package services

import (
	"context"
	"todo-list-be/internal/domain"
	"todo-list-be/internal/dto"
)


type TodoService interface{
	CreateTodo(request *dto.CreateTodoRequest) (domain.Todo, error)
	UpdateTodo(ctx context.Context, request *dto.UpdateTodoRequest) (domain.Todo, error)
}