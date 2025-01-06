package main

import (
	"cmp"
	"fmt"
	"slices"
)

type Product struct {
	name string
	price float64
	date string
}

func main() {

	products := []Product {
		{
			name: "samsung",
			price: 20,
			date: "2025-02-12",
		},
		{
			name: "oppo",
			price: 22,
			date: "2025-04-02",
		},
		{
			name: "redmi",
			price: 21,
			date: "2020-02-12",
		},
		{
			name: "nokia",
			price: 23,
			date: "2022-10-03",
		},
	}
	
	slices.SortFunc(products, func(a, b Product) int {
		return cmp.Compare(b.date, a.date)
	})

	numbers := []int {2, 4, 8, 10, 5, 7, 1}
	slices.Sort(numbers)

	fmt.Println(products)
	fmt.Println(numbers)
}
