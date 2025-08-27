package repository

import (
	"todo-list-be/internal/domain"
	"todo-list-be/internal/dto"

	"github.com/gofiber/fiber/v2"
)

type TodoRepository interface{
	Save(todo *dto.CreateTodoRequest) (domain.Todo, error)
	Update(ctx *fiber.Ctx, todo *dto.UpdateTodoRequest, id int) (domain.Todo, error)
	GetById(ctx *fiber.Ctx, id int) (domain.Todo, error)
}