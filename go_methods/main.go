package main

import (
	"fmt"
)

type Product struct {
		name   string
		price  float64
		inSale bool
	}

func main() {

	products := []Product{
		{
			name:   "samsung",
			price:  20.5,
			inSale: false,
		},
		{
			name:   "redmi",
			price:  25.5,
			inSale: true,
		},
		{
			name:   "nokia",
			price:  10.0,
			inSale: false,
		},
		{
			name:   "oppo",
			price:  22.3,
			inSale: true,
		},
	}

    productNames := mapSlice(products, func(p Product) string {
        return p.name
    })

    productDoublePrice := mapSlice(products, func(p Product) float64 {
        return p.price * 2
    })

    fmt.Println("Names:", productNames)
    fmt.Println("Double prices:", productDoublePrice)
}

// map 
func mapSlice[T any, R any](items []T, transform func(T) R) []R {
    newItems := make([]R, 0, len(items))
    for _, item := range items {
        newItems = append(newItems, transform(item))
    }
    return newItems
}


