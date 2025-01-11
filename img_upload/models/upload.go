package models

import "time"

type Image struct {
	Filename string `json:"filename"`
	FilePath string `json:"file_path"`
}

type Upload struct {
	Id          string    `json:"id" gorm:"primaryKey"`
	Image       Image     `json:"image" gorm:"embedded;embeddedPrefix:image_"`
	Description string    `json:"description"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedAt   time.Time `json:"created_at"`
}
