package main

import "time"

type Media struct {
	MediaFilename string `json:"media_filename"`
	MediaFilePath string `json:"media_file_path"`
}

type Variation struct {
	VariationType string `json:"variation_type"`
	VariationTag  string `json:"variation_tag"`
}

type Product struct {
	ID              string    `json:"id"`
	ThumbnailName   string    `json:"thumbnail_name"`
	ThumbnailPath   string    `json:"thumbnail_path"`
	Status          string    `json:"status"`
	Category        string    `json:"category"`
	Tag             string    `json:"tag"`
	Template        string    `json:"template"`
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	Price           float64   `json:"price"`
	DiscountType    string    `json:"discount_type"`
	TaxClass        string    `json:"tax_class"`
	VATAmount       float64   `json:"vat_amount"`
	SKUNumber       string    `json:"sku_number"`
	BarcodeNumber   string    `json:"barcode_number"`
	OnShelf         int64     `json:"on_shelf"`
	OnWarehouse     int64     `json:"on_warehouse"`
	AllowBackOrder  bool      `json:"allow_backorder"`
	InPhysical      bool      `json:"in_physical"`
	MetaTitle       string    `json:"meta_title"`
	MetaDescription string    `json:"meta_description"`
	MetaKeywords    string    `json:"meta_keywords"`
	CreatedAt       time.Time `json:"created_at"`
}

// new product
func NewProduct(id, thumbnailName, thumbnailPath, status, category, tag, template, name, description, discountType, taxClass, skuNumber, barcodeNumber, metaTitle, metaDescription, metaKeywords string, price, vatAmount float64, onShelf, onWarehouse int64, allowBackorder, inPhysical bool) (*Product, error) {
	return &Product{
		ID: id,
		ThumbnailName: thumbnailName,
		ThumbnailPath: thumbnailPath,
		Status: status,
		Category: category,
		Tag: tag,
		Template: template,
		Name: name,
		Description: description,
		Price: price,
		DiscountType: discountType,
		TaxClass: taxClass,
		VATAmount: vatAmount,
		SKUNumber: skuNumber,
		BarcodeNumber: barcodeNumber,
		OnShelf: onShelf,
		OnWarehouse: onWarehouse,
		AllowBackOrder: allowBackorder,
		InPhysical: inPhysical,
		MetaTitle: metaTitle,
		MetaDescription: metaDescription,
		MetaKeywords: metaKeywords,
		CreatedAt: time.Now().UTC(),
	}, nil
}