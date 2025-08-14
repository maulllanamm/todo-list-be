package dto

import "time"


type CreateTodoResponse struct {
    ID          int        `json:"id"`
    Title       string     `json:"title"`
    Description *string    `json:"description,omitempty"`
    Completed   bool       `json:"completed"`
    Priority    string     `json:"priority"`
    DueDate     *time.Time `json:"due_date,omitempty"`
    CreatedAt   time.Time  `json:"created_at"`
    UpdatedAt   time.Time  `json:"updated_at"`
}