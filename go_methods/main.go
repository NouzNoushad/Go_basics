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

	filterPrice := filter(products, func(p Product) bool {
		return p.price > 20
	})

	reducePrice := reduce(products, 0, func(acc float64, p Product) float64 {
		return acc + p.price
	})

	fmt.Println("Names:", productNames)
	fmt.Println("Double prices:", productDoublePrice)
	fmt.Println("Filter prices greater than 21:", filterPrice)
	fmt.Println("Reduce prices:", reducePrice)
}

// map
func mapSlice[T any, R any](items []T, transform func(T) R) []R {
	newItems := make([]R, 0, len(items))
	for _, item := range items {
		newItems = append(newItems, transform(item))
	}
	return newItems
}

// filter
func filter[T any](items []T, transform func(T) bool) []T {
	filteredItems := make([]T, 0, len(items))
	for _, item := range items {
		if transform(item) {
			filteredItems = append(filteredItems, item)
		}
	}
	return filteredItems
}

// reduce / fold
func reduce[T any](items []T, initial float64, accumulator func(float64, T) float64) float64 {
	result := initial
	for _, item := range items {
		result = accumulator(result, item)
	}
	return result
}
