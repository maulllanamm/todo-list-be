package services

import (
	"todo-list-be/internal/domain"
	"todo-list-be/internal/dto"

	"github.com/gofiber/fiber/v2"
)


type TodoService interface{
	CreateTodo(request *dto.CreateTodoRequest) (domain.Todo, error)
	UpdateTodo(ctx *fiber.Ctx, request *dto.UpdateTodoRequest, id int) (domain.Todo, error)
	GetTodoById(ctx *fiber.Ctx, id int) (domain.Todo, error)
}