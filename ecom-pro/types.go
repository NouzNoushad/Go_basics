package main

type Media struct {
	ID            string `json:"id"`
	ProductID     string `json:"product_id"`
	MediaFilename string `json:"media_filename"`
	MediaFilePath string `json:"media_file_path"`
	CreatedAt     string `json:"created_at"`
}

type Variation struct {
	ID            string `json:"id"`
	ProductID     string `json:"product_id"`
	VariationType string `json:"variation_type"`
	VariationTag  string `json:"variation_tag"`
	CreatedAt     string `json:"created_at"`
}

type Product struct {
	ID              string      `json:"id"`
	ThumbnailName   string      `json:"thumbnail_name"`
	ThumbnailPath   string      `json:"thumbnail_path"`
	Status          string      `json:"status"`
	Category        string      `json:"category"`
	Tag             string      `json:"tag"`
	Template        string      `json:"template"`
	Name            string      `json:"name"`
	Description     string      `json:"description"`
	Price           float64     `json:"price"`
	DiscountType    string      `json:"discount_type"`
	TaxClass        string      `json:"tax_class"`
	VATAmount       float64     `json:"vat_amount"`
	SKUNumber       string      `json:"sku_number"`
	BarcodeNumber   string      `json:"barcode_number"`
	OnShelf         int64       `json:"on_shelf"`
	OnWarehouse     int64       `json:"on_warehouse"`
	AllowBackOrder  bool        `json:"allow_backorder"`
	InPhysical      bool        `json:"in_physical"`
	MetaTitle       string      `json:"meta_title"`
	MetaDescription string      `json:"meta_description"`
	MetaKeywords    string      `json:"meta_keywords"`
	Variations      []Variation `json:"variations"`
	Media           []Media     `json:"media"`
	CreatedAt       string      `json:"created_at"`
}

// new product
func NewProduct(product *Product) (*Product, error) {
	return &Product{
		ID:              product.ID,
		ThumbnailName:   product.ThumbnailName,
		ThumbnailPath:   product.ThumbnailPath,
		Status:          product.Status,
		Category:        product.Category,
		Tag:             product.Tag,
		Template:        product.Template,
		Name:            product.Name,
		Description:     product.Description,
		Price:           product.Price,
		DiscountType:    product.DiscountType,
		TaxClass:        product.TaxClass,
		VATAmount:       product.VATAmount,
		SKUNumber:       product.SKUNumber,
		BarcodeNumber:   product.BarcodeNumber,
		OnShelf:         product.OnShelf,
		OnWarehouse:     product.OnWarehouse,
		AllowBackOrder:  product.AllowBackOrder,
		InPhysical:      product.InPhysical,
		MetaTitle:       product.MetaTitle,
		MetaDescription: product.MetaDescription,
		MetaKeywords:    product.MetaKeywords,
		CreatedAt:       product.CreatedAt,
	}, nil
}

// new media
func NewMedia(media *Media) (*Media, error) {
	return &Media{
		ID:            media.ID,
		ProductID:     media.ProductID,
		MediaFilename: media.MediaFilename,
		MediaFilePath: media.MediaFilePath,
		CreatedAt:     media.CreatedAt,
	}, nil
}

// new variation
func NewVariation(variation *Variation) (*Variation, error) {
	return &Variation{
		ID:            variation.ID,
		ProductID:     variation.ProductID,
		VariationType: variation.VariationType,
		VariationTag:  variation.VariationTag,
		CreatedAt:     variation.CreatedAt,
	}, nil
}
