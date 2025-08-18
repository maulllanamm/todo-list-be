package dto

import "time"

type UpdateTodoRequest struct{
    Id          int        `json:"id"`
	Title       string     `json:"title" validate:"required,min=1,max=255"`
    Description *string    `json:"description,omitempty"`
    Completed   bool       `json:"completed"`
    Priority    string     `json:"priority" validate:"oneof=low medium high"`
    DueDate     *time.Time `json:"due_date,omitempty"`
}