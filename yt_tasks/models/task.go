package models

import "time"

type Image struct {
	Filename string `json:"filename"`
	FilePath string `json:"file_path"`
}

type Assignee struct {
	Username string `json:"username"`
	Image    *Image  `json:"image,omitempty" gorm:"embedded;embeddedPrefix=image_"`
}

type Task struct {
	Id          string    `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Cateogry    string    `json:"category"`
	Priority    string    `json:"priority"`
	Status      string    `json:"status"`
	Assignee    Assignee  `json:"assignee" gorm:"embedded;embeddedPrefix=assignee_"`
	DueDate     time.Time `json:"due_date"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedAt   time.Time `json:"created_at"`
}
