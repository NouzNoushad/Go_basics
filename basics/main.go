package main

import (
	"encoding/json"
	"fmt"
)

type Product struct {
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
	InSale bool    `json:"inSale"`
}

func main() {
	// JSON
	products := []Product{
		{
			Name:   "samsung",
			Price:  20.5,
			InSale: false,
		},
		{
			Name:   "redmi",
			Price:  22.5,
			InSale: true,
		},
		{
			Name:   "nokia",
			Price:  25.5,
			InSale: true,
		},
		{
			Name:   "oppo",
			Price:  21.5,
			InSale: false,
		},
	}

	// convert data to json
	jsonData, err := json.Marshal(products)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	productsList := string(jsonData)
	fmt.Println(productsList)

	var p []Product

	// convert json to data
	error := json.Unmarshal([]byte(productsList), &p)
	if error != nil {
		fmt.Println("Error", error)
	}

	fmt.Println(p[3].Price)
}
