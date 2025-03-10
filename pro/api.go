package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type APIServer struct {
	listenAddr string
	store      Storage
}

func NewAPIServer(listenAddr string, store Storage) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		store:      store,
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/product", makeHandleFunc(s.handleProduct))
	router.HandleFunc("/product/{id}", makeHandleFunc(s.handleProductByID))

	http.ListenAndServe(s.listenAddr, router)
}

func (s *APIServer) handleProduct(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "POST" {
		return s.handleCreateProduct(w, r)
	}

	if r.Method == "GET" {
		return s.handleGetProducts(w, r)
	}

	return fmt.Errorf("method not allowed %s", r.Method)
}

func (s *APIServer) handleProductByID(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGetProductByID(w, r)
	}

	if r.Method == "DELETE" {
		return s.handleDeleteProduct(w, r)
	}

	if r.Method == "PUT" {
		return s.handleUpdateProduct(w, r)
	}

	return fmt.Errorf("method not allowed %s", r.Method)
}

func validationError(err string) error {
	return errors.New(err)
}

func validateProduct(req *Product) error {
	if req.Name == "" {
		return validationError("Name is required")
	}

	if req.Brand == "" {
		return validationError("Brand is required")
	}

	if !ValidCategory(req.Category) {
		return validationError("Invalid category")
	}

	if req.Price == 0 {
		return validationError("Price is required")
	}

	if req.Quantity == 0 {
		return validationError("Quantity is required")
	}

	return nil
}

func (s *APIServer) createProduct(req *Product) (*Product, error) {
	id := uuid.New().String()

	product, err := NewProduct(id, req.Name, req.Brand, req.Category, req.Price, req.Quantity)
	if err != nil {
		return nil, err
	}

	if err := s.store.CreateProduct(product); err != nil {
		return nil, err
	}

	return product, nil
}

func (s *APIServer) handleCreateProduct(w http.ResponseWriter, r *http.Request) error {
	req := new(Product)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		return WriteJSON(w, http.StatusBadRequest, APIError{Error: "Invalid JSON format"})
	}

	if err := validateProduct(req); err != nil {
		return WriteJSON(w, http.StatusBadRequest, APIError{Error: err.Error()})
	}

	product, err := s.createProduct(req)
	if err != nil {
		return WriteJSON(w, http.StatusInternalServerError, APIError{Error: err.Error()})
	}

	// success
	return WriteJSON(w, http.StatusCreated, map[string]interface{}{
		"message": "Product created",
		"data":    product,
	})
}

func (s *APIServer) handleGetProducts(w http.ResponseWriter, _ *http.Request) error {

	products, err := s.store.GetProducts()
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, map[string]interface{}{
		"data":  products,
		"items": fmt.Sprintf("%d items", len(products)),
	})
}

func (s *APIServer) handleGetProductByID(w http.ResponseWriter, r *http.Request) error {

	id := getID(r)

	product, err := s.store.GetProductByID(id)
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, map[string]interface{}{
		"data": product,
	})
}

func (s *APIServer) handleDeleteProduct(w http.ResponseWriter, r *http.Request) error {

	id := getID(r)

	if err := s.store.DeleteProduct(id); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, map[string]interface{}{
		"message": "product deleted",
		"id":      id,
	})
}

func (s *APIServer) handleUpdateProduct(w http.ResponseWriter, r *http.Request) error {

	id := getID(r)

	extProduct, err := s.store.GetProductByID(id)
	if err != nil {
		return err
	}

	req := new(Product)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		return err
	}

	product, err := UpdateProduct(extProduct, req.Name, req.Brand, req.Category, req.Price, req.Quantity)
	if err != nil {
		return err
	}

	if err := s.store.UpdateProduct(id, product); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, map[string]interface{}{
		"message": "Product updated",
		"data":    product,
	})
}

func getID(r *http.Request) string {
	id := mux.Vars(r)["id"]

	return id
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}

type apiFunc func(http.ResponseWriter, *http.Request) error

type APIError struct {
	Error string `json:"error"`
}

func makeHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, APIError{Error: err.Error()})
		}
	}
}
