package repository

import (
	"todo-list-be/internal/domain"
	"todo-list-be/internal/dto"
)

type TodoRepository interface{
	Save(todo *dto.CreateTodoRequest) (domain.Todo, error)
}