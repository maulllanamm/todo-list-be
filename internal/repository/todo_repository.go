package repository

import (
	"context"
	"todo-list-be/internal/domain"
	"todo-list-be/internal/dto"
)

type TodoRepository interface{
	Save(todo *dto.CreateTodoRequest) (domain.Todo, error)
	Update(ctx context.Context, todo *dto.UpdateTodoRequest) (domain.Todo, error)
}