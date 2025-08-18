package repository

import (
	"context"
	"database/sql"
	"fmt"
	"todo-list-be/internal/domain"
	"todo-list-be/internal/dto"
)

type TodoRepositoryImplementation struct{
	db *sql.DB
}

func NewTodoRepository(db *sql.DB) TodoRepository{
	return &TodoRepositoryImplementation{db: db}
}

func (r *TodoRepositoryImplementation) Save(todo *dto.CreateTodoRequest) (domain.Todo, error) {
	query := `
		INSERT INTO todos (title, description, priority, due_date)
		VALUES ($1, $2, $3, $4)
		RETURNING id, title, description, completed, priority, due_date, created_at, updated_at
	`

	var result domain.Todo
	err := r.db.QueryRow(query, todo.Title, todo.Description, todo.Priority, todo.DueDate).Scan(
		&result.Id,
		&result.Title,
		&result.Description,
		&result.Completed,
		&result.Priority,
		&result.DueDate,
		&result.CreatedAt,
		&result.UpdatedAt,
	)

	if err != nil {
		return domain.Todo{}, fmt.Errorf("failed to create todo: %w", err)
	}

	return result, nil
}

func (r *TodoRepositoryImplementation) Update(ctx context.Context, todo *dto.UpdateTodoRequest) (domain.Todo, error) {
	query := `
		UPDATE todos
		SET title = $1,
			description = $2,
			completed = $3,
			priority = $4,
			due_date = $5,
			updated_at = NOW()
		WHERE id = $6
	`

	var result domain.Todo
	_, err := r.db.ExecContext(ctx, query,
		todo.Title,
		todo.Description,
		todo.Completed,
		todo.Priority,
		todo.DueDate,
		todo.Id,
	)

	if err != nil {
		return domain.Todo{}, fmt.Errorf("failed to update todo: %w", err)
	}

	return result, nil
}