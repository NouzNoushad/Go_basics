package main

import (
	"time"
)

type Product struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Brand     string    `json:"brand"`
	Category  string    `json:"category"`
	Price     float64   `json:"price"`
	Quantity  int64     `json:"quantity"`
	CreatedAt time.Time `json:"created_at"`
}

func NewProduct(id, name, brand, category string, price float64, quantity int64) (*Product, error) {
	return &Product{
		ID:        id,
		Name:      name,
		Brand:     brand,
		Category:  category,
		Price:     price,
		Quantity:  quantity,
		CreatedAt: time.Now().UTC(),
	}, nil
}

func ValidCategory(category string) bool {
	switch category {
	case "fashion",
		"food",
		"electronics",
		"shirt",
		"furniture",
		"toys",
		"cosmetics":
		return true
	}
	return false
}
