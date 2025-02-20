package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

// Interface
type Storage interface {
	AddProduct(*Product) error
	EditProduct(int, *Product) error
	GetProducts() ([]*Product, error)
	GetProductByID(int) (*Product, error)
	DeleteProduct(int) error
}

// Postgresql store
type PostgresStore struct {
	db *sql.DB
}

// Set up database
func NewPostgresStore() (*PostgresStore, error) {
	connStr := "host=localhost user=postgres password=noushad dbname=ecom-pro port=5432 sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStore{
		db: db,
	}, nil
}

// Init database
func (s *PostgresStore) InitDB() error {
	if err := s.createProductTable(); err != nil {
		return err
	}

	if err := s.createMediaTable(); err != nil {
		return err
	}

	return s.createVariationTable()
}

// Set up product table
func (s *PostgresStore) createProductTable() error {
	query := `create table if not exists product(
		id text primary key,
		thumbnail_name text,
		thumbnail_path text,
		status text,
		category text,
		tag text,
		template text,
		name text not null,
		description text,
		price numeric not null,
		discount_type text,
		tax_class text,
		vat_amount numeric,
		sku_number text not null,
		barcode_number text not null,
		on_shelf numeric not null,
		on_warehouse numeric,
		allow_backorder boolean not null,
		in_physical boolean,
		meta_title text,
		meta_description text,
		meta_keywords text,
		created_at timestamp default now()
	)`

	_, err := s.db.Exec(query)

	return err
}

// Set up media table
func (s *PostgresStore) createMediaTable() error {
	query := `create table if not exists media(
		id text primary key,
		product_id text references product(id) on delete cascade,
		media_filename text,
		media_file_path text,
		created_at timestamp default now()
	)`

	_, err := s.db.Exec(query)

	return err
}

// Set up media table
func (s *PostgresStore) createVariationTable() error {
	query := `create table if not exists variation(
		id text primary key,
		product_id text references product(id) on delete cascade,
		variation_type text,
		variation_tag text,
		created_at timestamp default now()
	)`

	_, err := s.db.Exec(query)

	return err
}

// Add Product
func (s *PostgresStore) AddProduct(product *Product) error {
	query := `insert into product (
		id,
		thumbnail_name,
		thumbnail_path,
		status,
		category,
		tag,
		template,
		name,
		description,
		price,
		discount_type,
		tax_class,
		vat_amount,
		sku_number,
		barcode_number,
		on_shelf,
		on_warehouse,
		allow_backorder,
		in_physical,
		meta_title,
		meta_description,
		meta_keywords,
		created_at) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23)`
	_, err := s.db.Query(
		query,
		product.ID,
		product.ThumbnailName,
		product.ThumbnailPath,
		product.Status,
		product.Category,
		product.Tag,
		product.Template,
		product.Name,
		product.Description,
		product.Price,
		product.DiscountType,
		product.TaxClass,
		product.VATAmount,
		product.SKUNumber,
		product.BarcodeNumber,
		product.OnShelf,
		product.OnWarehouse,
		product.AllowBackOrder,
		product.InPhysical,
		product.MetaTitle,
		product.MetaDescription,
		product.MetaKeywords,
		product.CreatedAt,
	)

	return err
}

// Edit Product
func (s *PostgresStore) EditProduct(id int, product *Product) error {
	return nil
}

// Get Products
func (s *PostgresStore) GetProducts() ([]*Product, error) {
	return nil, nil
}

// Get Product by id
func (s *PostgresStore) GetProductByID(id int) (*Product, error) {
	return nil, nil
}

// Delete Product
func (s *PostgresStore) DeleteProduct(id int) error {
	return nil
}
