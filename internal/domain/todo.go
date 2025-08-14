package domain

import "time"


type Todo struct {
	Id          int
	Title       string     
	Description *string
	Completed   bool
	Priority    string    
	DueDate     *time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
