package main

import (
	"todo-list-be/internal/database"
	"todo-list-be/internal/handler"
	"todo-list-be/internal/repository"
	"todo-list-be/internal/services"

	"github.com/gofiber/fiber/v2"
)

func main(){
	db := database.NewConnection()
	
	// initialize repo
	todoRepository := repository.NewTodoRepository(db)

	// initialize service
	todoService := services.NewTodoService(todoRepository)

	// initialize handler
	todoHandler := handler.NewTodoHandler(todoService)

	app := fiber.New()

	// API v1 routes
	api := app.Group("/api/v1")
	
	// Todo routes
	todos := api.Group("/todos")
	todos.Post("/", todoHandler.CreateTodo)

	app.Listen(":3000")
}