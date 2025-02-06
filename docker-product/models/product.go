package models

import "time"

type Product struct {
	Id          string    `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name"`
	Brand       string    `json:"brand"`
	Category    string    `json:"category"`
	Price       float32   `json:"price"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}
