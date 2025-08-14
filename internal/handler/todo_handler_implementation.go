package handler

import (
	"todo-list-be/internal/dto"
	"todo-list-be/internal/services"

	"github.com/gofiber/fiber/v2"
)

type TodoHandlerImplementation struct{
	service services.TodoService
}

func NewTodoHandler(service services.TodoService) TodoHandler{
	return &TodoHandlerImplementation{service: service}
}

func(s *TodoHandlerImplementation) CreateTodo(c *fiber.Ctx) error{
	var request dto.CreateTodoRequest

	if err := c.BodyParser(&request); err != nil {
		msg := err.Error()
		response := dto.APIResponse{
			Success: false,
			Message: "Failed to parse request body",
			Error:   &msg,
		}
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	todo := dto.CreateTodoRequest{
		Title:       request.Title,
		Description: request.Description,
		Completed: 	 request.Completed,
		Priority:    request.Priority,
		DueDate:     request.DueDate,
	}

	created, err := s.service.CreateTodo(&todo)
	if err != nil {
		msg := err.Error()
		response := dto.APIResponse{
			Success: false,
			Message: "Failed to create todo",
			Error:   &msg,
		}
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	todoResponse := dto.CreateTodoResponse{
		ID:          created.Id,
		Title:       created.Title,
		Description: created.Description,
		Completed:   created.Completed,
		Priority:    created.Priority,
		DueDate:     created.DueDate,
		CreatedAt:   created.CreatedAt,
		UpdatedAt:   created.UpdatedAt,
	}

	response := dto.APIResponse{
		Success: true,
		Message: "Todo created successfully",
		Data:    todoResponse,
	}

	return c.Status(fiber.StatusCreated).JSON(response)
}