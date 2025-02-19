package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type APIServer struct {
	listenAddr string
	store      Storage
}

func NewAPIServer(listenAdr string, store Storage) *APIServer {
	return &APIServer{
		listenAddr: listenAdr,
		store:      store,
	}
}

// router
func (s *APIServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/product", makeHandleFunc(s.handleProduct))

	http.ListenAndServe(s.listenAddr, router)
}

// handle request methods (product)
func (s *APIServer) handleProduct(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGetProducts(w, r)
	}

	if r.Method == "POST" {
		return s.handleAddProduct(w, r)
	}

	return fmt.Errorf("method not allowed %s", r.Method)
}

// handle add product
func (s *APIServer) handleAddProduct(w http.ResponseWriter, r *http.Request) error {
	// parse multipart form
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		return WriteJSON(w, http.StatusBadRequest, ApiError{Error: "Failed to parse form"})
	}

	// id
	id := uuid.New().String()

	// thumbnail
	file, fileHeader, err := r.FormFile("thumbnail")
	if err != nil {
		return WriteJSON(w, http.StatusInternalServerError, ApiError{Error: "Thumbnail upload failed"})
	}
	defer file.Close()

	// create upload dir
	uploadDir := "uploads"
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		return WriteJSON(w, http.StatusInternalServerError, ApiError{Error: "Failed to create uploads directory"})
	}

	// save file
	fileName := uuid.New().String() + fileHeader.Filename
	filePath := filepath.Join(uploadDir, fileName)
	outFile, err := os.Create(filePath)
	if err != nil {
		return WriteJSON(w, http.StatusInternalServerError, ApiError{Error: "Failed to save file"})
	}
	defer outFile.Close()

	// copy file contents
	_, err = io.Copy(outFile, file)
	if err != nil {
		return WriteJSON(w, http.StatusInternalServerError, ApiError{Error: "Failed to copy file"})
	}

	// 		status,
	status := r.FormValue("status")

	// 		category,
	category := r.FormValue("category")

	// 		tag,
	tag := r.FormValue("tag")

	// 		template,
	template := r.FormValue("template")

	// 		name,
	name := r.FormValue("name")
	if name == "" {
		return WriteJSON(w, http.StatusBadRequest, ApiError{Error: "Name is required"})
	}

	// 		description,
	description := r.FormValue("description")

	// 		price,
	price := r.FormValue("price")
	priceParsed, err := strconv.ParseFloat(price, 64)
	if err != nil {
		return WriteJSON(w, http.StatusBadRequest, ApiError{Error: "Invalid price format"})
	}
	if priceParsed == 0 {
		return WriteJSON(w, http.StatusBadRequest, ApiError{Error: "Price is required"})
	}

	// 		discount_type,
	discountType := r.FormValue("discount_type")

	// 		tax_class,
	taxClass := r.FormValue("tax_class")
	if taxClass == "" {
		return WriteJSON(w, http.StatusBadRequest, ApiError{Error: "Tax class is required"})
	}

	// 		vat_amount,
	vatAmount := r.FormValue("vat_amount")
	vatAmountParsed, err := strconv.ParseFloat(vatAmount, 64)
	if err != nil {
		return WriteJSON(w, http.StatusBadRequest, ApiError{Error: "Invalid VAT Amount format"})
	}

	// 		sku_number,
	skuNumber := r.FormValue("sku_number")
	if skuNumber == "" {
		return WriteJSON(w, http.StatusBadRequest, ApiError{Error: "SKU Number is required"})
	}

	// 		barcode_number,
	barcodeNumber := r.FormValue("barcode_number")
	if barcodeNumber == "" {
		return WriteJSON(w, http.StatusBadRequest, ApiError{Error: "Barcode Number is required"})
	}

	// 		on_shelf,
	onShelf := r.FormValue("on_shelf")
	onShelfParsed, err := strconv.ParseInt(onShelf, 0, 64)
	if err != nil {
		return WriteJSON(w, http.StatusBadRequest, ApiError{Error: "Invalid on shelf format"})
	}
	if onShelfParsed == 0 {
		return WriteJSON(w, http.StatusBadRequest, ApiError{Error: "Quantity on shelf is required"})
	}

	// 		on_warehouse,
	onWarehouse := r.FormValue("on_warehouse")
	onWarehouseParsed, err := strconv.ParseInt(onWarehouse, 0, 64)
	if err != nil {
		return WriteJSON(w, http.StatusBadRequest, ApiError{Error: "Invalid on warehouse format"})
	}

	// 		allow_backorder,
	allowBackorder := r.FormValue("allow_backorder")
	allowBackorderParsed, err := strconv.ParseBool(allowBackorder)
	if err != nil {
		return WriteJSON(w, http.StatusBadRequest, ApiError{Error: "Invalid allow backorder format"})
	}

	// 		in_physical,
	inPhysical := r.FormValue("in_physical")
	inPhysicalParsed, err := strconv.ParseBool(inPhysical)
	if err != nil {
		return WriteJSON(w, http.StatusBadRequest, ApiError{Error: "Invalid in physical format"})
	}

	// 		meta_title,
	metaTitle := r.FormValue("meta_title")

	// 		meta_description,
	metaDescription := r.FormValue("meta_description")
	
	// 		meta_keywords,
	metaKeywords := r.FormValue("meta_keywords")

	product, err := NewProduct(id, fileName, filePath, status, category, tag, template, name, description, discountType, taxClass, skuNumber, barcodeNumber, metaTitle, metaDescription, metaKeywords, priceParsed, vatAmountParsed, onShelfParsed, onWarehouseParsed, allowBackorderParsed, inPhysicalParsed)
	if err != nil {
		return WriteJSON(w, http.StatusInternalServerError, ApiError{Error: err.Error()})
	}

	if err := s.store.AddProduct(product); err != nil {
		return WriteJSON(w, http.StatusInternalServerError, ApiError{Error: err.Error()})
	}

	return WriteJSON(w, http.StatusCreated, map[string]interface{}{
		"message": "Product added",
		"data":    product,
	})
}

// handle get product
func (s *APIServer) handleGetProducts(w http.ResponseWriter, _ *http.Request) error {
	return WriteJSON(w, http.StatusOK, map[string]string{
		"message": "get product",
	})
}

// json output
func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}

// api error
type ApiError struct {
	Error string `json:"error"`
}

type apiFunc func(http.ResponseWriter, *http.Request) error

// handle func wrapper
func makeHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}
