package handler

import (
	"github.com/gofiber/fiber/v2"
)

type TodoHandler interface{
	CreateTodo(ctx *fiber.Ctx) error
	UpdateTodo(ctx *fiber.Ctx) error
	GetTodoById(ctx *fiber.Ctx) error
}