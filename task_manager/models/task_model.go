package models

import "time"

type Task struct {
	Id          string    `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	Category    string    `json:"category"`
	StartDate   time.Time `json:"start_date"`
	DueDate     time.Time `json:"due_date"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedAt   time.Time `json:"created_at"`
}
