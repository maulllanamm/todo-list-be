package services

import (
	"todo-list-be/internal/domain"
	"todo-list-be/internal/dto"
)


type TodoService interface{
	CreateTodo(request *dto.CreateTodoRequest) (domain.Todo, error)
}