package utils

import (
	"todo-list-be/internal/dto"

	"github.com/gofiber/fiber/v2"
)


func SuccessResponse(c *fiber.Ctx, status int, message string, data interface{}) error {
	response := dto.APIResponse{
		Success: true,
		Message: message,
		Data:    data,
	}
	return c.Status(status).JSON(response)
}

func ErrorResponse(c *fiber.Ctx, status int, message string, err string) error {
	response := dto.APIResponse{
		Success: false,
		Message: message,
		Error:   &err,
	}
	return c.Status(status).JSON(response)
}

func ValidationErrorResponse(c *fiber.Ctx, errors []string) error {
	errorStr := "Validation failed"
	if len(errors) > 0 {
		errorStr = errors[0]
	}
	
	response := dto.APIResponse{
		Success: false,
		Message: "Validation failed",
		Error:   &errorStr,
		Data:    errors,
	}
	return c.Status(fiber.StatusBadRequest).JSON(response)
}