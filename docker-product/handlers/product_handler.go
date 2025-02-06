package handlers

import (
	"docker-product/database"
	"docker-product/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func isEmptyField(field string) bool {
	return field == ""
}

// Create product
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product map[string]interface{}

	json.NewDecoder(r.Body).Decode(&product)

	id := uuid.New().String()

	name, ok := product["name"].(string)
	if !ok || isEmptyField(name) {
		http.Error(w, "Name cannot be empty", http.StatusBadRequest)
		return
	}

	brand, ok := product["brand"].(string)
	if !ok || isEmptyField(brand) {
		http.Error(w, "Brand cannot be empty", http.StatusBadRequest)
		return
	}

	category, ok := product["category"].(string)
	if !ok || isEmptyField(category) {
		http.Error(w, "Category cannot be empty", http.StatusBadRequest)
		return
	}

	var price float64
	switch vPrice := product["price"].(type) {
	case float64:
		price = vPrice
	case string:
		priceParsed, err := strconv.ParseFloat(vPrice, 64)
		if err != nil {
			http.Error(w, "Price must be valid number", http.StatusBadRequest)
			return
		}
		price = priceParsed
	default:
		http.Error(w, "Price must be valid number", http.StatusBadRequest)
		return
	}

	description, _ := product["description"].(string)

	newProduct := models.Product{
		Id:          id,
		Name:        name,
		Brand:       brand,
		Category:    category,
		Price:       float32(price),
		Description: description,
	}

	if err := database.DB.Create(&newProduct).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "product created",
		"product": product})
}

// get products
func GetProducts(w http.ResponseWriter, r *http.Request) {
	var products []models.Product

	if err := database.DB.Order("created_at desc").Find(&products).Error; err != nil {
		http.Error(w, "Failed to fetch products", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"data":  products,
		"items": fmt.Sprintf("%d items", len(products))})
}

// get product by id
func GetProductDetails(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	params := mux.Vars(r)
	id := params["id"]

	if err := database.DB.Where("id = ?", id).First(&product).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{"data": product})
}

// delete product
func DeleteProductDetails(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	params := mux.Vars(r)
	id := params["id"]

	if err := database.DB.Where("id = ?", id).First(&product).Error; err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := database.DB.Delete(&product).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{"message": "Product deleted"})
}

// update product
func UpdateProductDetails(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	// existing products
	var product models.Product
	if err := database.DB.Where("id = ?", id).First(&product).Error; err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// request body
	var update map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// name
	if name, ok := update["name"].(string); ok {
		if !isEmptyField(name) {
			product.Name = name
		}
	}

	// brand
	if brand, ok := update["brand"].(string); ok {
		if !isEmptyField(brand) {
			product.Brand = brand
		}
	}

	// category
	if category, ok := update["category"].(string); ok {
		if !isEmptyField(category) {
			product.Category = category
		}
	}

	// price
	if price, ok := update["price"]; ok {
		if price != 0.0 {
			var vPrice float64
			switch v := price.(type) {
			case float64:
				vPrice = v
			case string:
				priceParsed, err := strconv.ParseFloat(v, 64)
				if err != nil {
					http.Error(w, "Price must be a valid number", http.StatusBadRequest)
					return
				}
				vPrice = priceParsed
			default:
				http.Error(w, "Price must be a valid number", http.StatusBadRequest)
				return
			}
			product.Price = float32(vPrice)
		}
	}

	// description
	if description, ok := update["description"].(string); ok {
		if !isEmptyField(description) {
			product.Description = description
		}
	}

	if err := database.DB.Save(&product).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Product updated",
		"data":    product,
	})
}
