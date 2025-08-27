package services

import (
	"todo-list-be/internal/domain"
	"todo-list-be/internal/dto"
	"todo-list-be/internal/repository"

	"github.com/gofiber/fiber/v2"
)

type TodoServiceImplementation struct{
	repo repository.TodoRepository
}

func NewTodoService(repo repository.TodoRepository) TodoService{
	return &TodoServiceImplementation{repo: repo}
}

func (s *TodoServiceImplementation) CreateTodo(request *dto.CreateTodoRequest) (domain.Todo, error){
	if request.Priority == "" {
		request.Priority = "Medium"
	}

	return s.repo.Save(request)
}

func (s *TodoServiceImplementation) UpdateTodo(ctx *fiber.Ctx, request *dto.UpdateTodoRequest, id int) (domain.Todo, error){
	return s.repo.Update(ctx, request, id)
}

func (s *TodoServiceImplementation) GetTodoById(ctx *fiber.Ctx, id int) (domain.Todo, error){
	return s.repo.GetById(ctx, id)
}