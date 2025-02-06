package models

import "time"

type User struct {
	Id        string    `json:"id" gorm:"primaryKey"`
	Username  string    `json:"username"`
	Email     string    `json:"email" gorm:"unique"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}
