package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

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
	router.HandleFunc("/product/{id}", makeHandleFunc(s.handleProductByID))
	router.HandleFunc("/media", makeHandleFunc(s.handleMedia))
	router.HandleFunc("/media/{id}", makeHandleFunc(s.handleMediaByID))

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

// handle request methods (product by id)
func (s *APIServer) handleProductByID(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGetProductByID(w, r)
	}

	if r.Method == "DELETE" {
		return s.handleDeleteProduct(w, r)
	}

	return fmt.Errorf("method not allowed %s", r.Method)
}

// handle request methods (media)
func (s *APIServer) handleMedia(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGetMedias(w, r)
	}

	if r.Method == "POST" {
		return s.handleAddMedia(w, r)
	}

	return fmt.Errorf("method not allowed %s", r.Method)
}

// handle request methods (media by id)
func (s *APIServer) handleMediaByID(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGetMediaByID(w, r)
	}

	return fmt.Errorf("method not allowed %s", r.Method)
}

// validation error
func validationError(err string) error {
	return errors.New(err)
}

// product validation
func productValidation(product *Product) error {
	// status,
	if !isValidStatus(product.Status) {
		return validationError("Invalid status")
	}

	// category,
	if !isValidCategory(product.Category) {
		return validationError("Invalid category")
	}

	// template,
	if !isValidTemplate(product.Template) {
		return validationError("Invalid template")
	}

	// name,
	if product.Name == "" {
		return validationError("Name is required")
	}

	// price,
	if product.Price == 0 {
		return validationError("Price is required and should not be zero")
	}

	// discount_type,
	if !isValidDiscountType(product.DiscountType) {
		return validationError("Invalid discount type")
	}

	// tax_class,
	if product.TaxClass == "" {
		return validationError("Tax class is required")
	}
	if !isValidTaxClass(product.TaxClass) {
		return validationError("Invalid tax class")
	}

	// sku_number,
	if product.SKUNumber == "" {
		return validationError("SKU Number is required")
	}

	// barcode_number,
	if product.BarcodeNumber == "" {
		return validationError("Barcode Number is required")
	}

	// on_shelf,
	if product.OnShelf == 0 {
		return validationError("Quantity on shelf is required")
	}

	return nil
}

// media validation
func mediaValidation(media *Media) error {

	// product id,
	if media.ProductID == "" {
		return validationError("Product ID is required")
	}

	return nil
}

// create file
func createFile(r *http.Request, rName string, dirName string) (string, string, error) {
	// thumbnail
	file, fileHeader, err := r.FormFile(rName)
	if err != nil {
		return "", "", validationError("File upload failed")
	}
	defer file.Close()

	// create upload dir
	uploadDir := dirName
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		return "", "", validationError("Failed to create uploads directory")
	}

	// save file
	fileName := uuid.New().String() + "_" + fileHeader.Filename
	filePath := filepath.Join(uploadDir, fileName)
	outFile, err := os.Create(filePath)
	if err != nil {
		return "", "", validationError("Failed to save file")
	}
	defer outFile.Close()

	// copy file contents
	_, err = io.Copy(outFile, file)
	if err != nil {
		return "", "", validationError("Failed to copy file")
	}

	return fileName, filePath, nil
}

// handle add product
func (s *APIServer) handleAddProduct(w http.ResponseWriter, r *http.Request) error {
	product := new(Product)

	// parse multipart form
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		return WriteJSON(w, http.StatusBadRequest, ApiError{Error: "Failed to parse form"})
	}

	product.ID = uuid.New().String()
	product.Status = r.FormValue("status")
	product.Category = r.FormValue("category")
	product.Tag = r.FormValue("tag")
	product.Template = r.FormValue("template")
	product.Name = r.FormValue("name")
	product.Description = r.FormValue("description")
	product.DiscountType = r.FormValue("discount_type")
	product.TaxClass = r.FormValue("tax_class")
	product.SKUNumber = r.FormValue("sku_number")
	product.BarcodeNumber = r.FormValue("barcode_number")
	product.MetaTitle = r.FormValue("meta_title")
	product.MetaDescription = r.FormValue("meta_description")
	product.MetaKeywords = r.FormValue("meta_keywords")

	product.Price, err = stringToFloat(r.FormValue("price"))
	if err != nil {
		return parseError(w, "Invalid price format")
	}
	product.VATAmount, err = stringToFloat(r.FormValue("vat_amount"))
	if err != nil {
		return parseError(w, "Invalid vat amount format")
	}

	product.OnShelf, err = stringToInt(r.FormValue("on_shelf"))
	if err != nil {
		return parseError(w, "Invalid on shelf format")
	}
	product.OnWarehouse, err = stringToInt(r.FormValue("on_warehouse"))
	if err != nil {
		return parseError(w, "Invalid on warehouse format")
	}

	product.AllowBackOrder, err = stringToBool(r.FormValue("allow_backorder"))
	if err != nil {
		return parseError(w, "Invalid allow backorder format")
	}
	product.InPhysical, err = stringToBool(r.FormValue("in_physical"))
	if err != nil {
		return parseError(w, "Invalide in physical format")
	}

	product.CreatedAt = time.Now().UTC().Format(time.RFC3339)

	// validation
	if err := productValidation(product); err != nil {
		return WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
	}

	// create file
	product.ThumbnailName, product.ThumbnailPath, err = createFile(r, "thumbnail", "uploads")
	if err != nil {
		return WriteJSON(w, http.StatusInternalServerError, ApiError{Error: err.Error()})
	}

	// variation model
	attributes := map[string][]string{
		"color":    splitValues(r.FormValue("color")),
		"size":     splitValues(r.FormValue("size")),
		"material": splitValues(r.FormValue("material")),
		"style":    splitValues(r.FormValue("style")),
	}

	// convert attr map to json
	variationJson, err := json.Marshal(attributes)
	if err != nil {
		return WriteJSON(w, http.StatusInternalServerError, err.Error())
	}

	// medias
	files := r.MultipartForm.File["media_files"]

	uploadedMedias := []*Media{}

	// create media dir
	mediaDir := "medias"
	if err := os.MkdirAll(mediaDir, os.ModePerm); err != nil {
		return WriteJSON(w, http.StatusInternalServerError, "Failed to create media directory")
	}

	for _, fileHeader := range files {
		media := new(Media)
		media.ID = uuid.New().String()
		media.ProductID = product.ID
		media.CreatedAt = time.Now().UTC().Format(time.RFC3339)

		file, err := fileHeader.Open()
		if err != nil {
			return WriteJSON(w, http.StatusInternalServerError, "Failed to open file")
		}
		defer file.Close()

		// save file
		fileName := uuid.New().String() + "_" + fileHeader.Filename
		filePath := filepath.Join(mediaDir, fileName)

		outFile, err := os.Create(filePath)
		if err != nil {
			return WriteJSON(w, http.StatusInternalServerError, "Failed to save file")
		}
		defer outFile.Close()

		// copy file contents
		_, err = io.Copy(outFile, file)
		if err != nil {
			return WriteJSON(w, http.StatusInternalServerError, "Failed to copy file")
		}

		media.MediaFilename = fileName
		media.MediaFilePath = filePath

		uploadedMedias = append(uploadedMedias, media)
	}

	product.Media = uploadedMedias

	// store
	if err := s.store.AddProductTransaction(product, variationJson, uploadedMedias); err != nil {
		return WriteJSON(w, http.StatusInternalServerError, ApiError{Error: err.Error()})
	}

	// success
	return WriteJSON(w, http.StatusCreated, map[string]interface{}{
		"message": "Product added",
		"data":    product,
	})
}

// handle get product
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

// handle get product by ID
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

// handle delete product
func (s *APIServer) handleDeleteProduct(w http.ResponseWriter, r *http.Request) error {

	id := getID(r)

    product, err := s.store.GetProductByID(id)
	if err != nil {
		return err
	}

    medias, err := s.store.GetMediasByProductID(id)
    if err != nil {
        return err
    }

    // remove thumbnail from uploads
    if err := os.Remove(product.ThumbnailPath); err != nil {
        return err
    }

    // remove images from medias
    for _, media := range medias {
        if err := os.Remove(media.MediaFilePath); err != nil {
            return err
        }
    }

	if err := s.store.DeleteProduct(id); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, map[string]interface{}{
		"message": "product deleted",
		"id":      id,
	})
}

// handle add media
func (s *APIServer) handleAddMedia(w http.ResponseWriter, r *http.Request) error {
	media := new(Media)

	// parse multipart form
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		return WriteJSON(w, http.StatusBadRequest, ApiError{Error: "Failed to parse form"})
	}

	media.ID = uuid.New().String()
	media.ProductID = r.FormValue("product_id")

	media.CreatedAt = time.Now().UTC().Format(time.RFC3339)

	// validation
	if err := mediaValidation(media); err != nil {
		return WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
	}

	// create file
	media.MediaFilename, media.MediaFilePath, err = createFile(r, "media_image", "medias")
	if err != nil {
		return WriteJSON(w, http.StatusInternalServerError, ApiError{Error: err.Error()})
	}

	// store
	if err := s.store.AddMedia(media); err != nil {
		return WriteJSON(w, http.StatusInternalServerError, ApiError{Error: err.Error()})
	}

	// success
	return WriteJSON(w, http.StatusCreated, map[string]interface{}{
		"message": "Media added",
		"data":    media,
	})
}

// handle get medias
func (s *APIServer) handleGetMedias(w http.ResponseWriter, _ *http.Request) error {

	medias, err := s.store.GetMedias()
	if err != nil {
		return err
	}

	// success
	return WriteJSON(w, http.StatusOK, map[string]interface{}{
		"data":  medias,
		"items": fmt.Sprintf("%d items", len(medias)),
	})
}

// handle get media by id
func (s *APIServer) handleGetMediaByID(w http.ResponseWriter, r *http.Request) error {

	id := getID(r)

	media, err := s.store.GetMediaByID(id)
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, map[string]interface{}{
		"data": media,
	})
}

// split values
func splitValues(value string) []string {
	if value == "" {
		return []string{}
	}
	values := strings.Split(value, ",")
	for i, v := range values {
		values[i] = strings.TrimSpace(v)
	}

	return values
}

// get ID
func getID(r *http.Request) string {
	id := mux.Vars(r)["id"]

	return id
}

// parse error
func parseError(w http.ResponseWriter, errStr string) error {
	return WriteJSON(w, http.StatusBadRequest, ApiError{Error: errStr})
}

// string to float
func stringToFloat(value string) (float64, error) {
	valueParsed, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0, err
	}
	return valueParsed, nil
}

// string to int
func stringToInt(value string) (int64, error) {
	valueParsed, err := strconv.ParseInt(value, 0, 64)
	if err != nil {
		return 0, err
	}
	return valueParsed, nil
}

// string to bool
func stringToBool(value string) (bool, error) {
	valueParsed, err := strconv.ParseBool(value)
	if err != nil {
		return false, err
	}
	return valueParsed, nil
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
