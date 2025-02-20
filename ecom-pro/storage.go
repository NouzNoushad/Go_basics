package main

import (
	"database/sql"
	"encoding/json"
	"time"

	_ "github.com/lib/pq"
)

// Interface
type Storage interface {
	AddProduct(*Product) error
	EditProduct(int, *Product) error
	GetProducts() ([]*Product, error)
	GetProductByID(int) (*Product, error)
	DeleteProduct(int) error
	AddMedia(*Media) error
	GetMedias() ([]*Media, error)
	AddVariation(*Variation) error
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

// Add Media
func (s *PostgresStore) AddMedia(media *Media) error {
	query := `insert into media (
		id, 
		product_id, 
		media_filename, 
		media_file_path, 
		created_at) values ($1, $2, $3, $4, $5)`

	_, err := s.db.Query(
		query,
		media.ID,
		media.ProductID,
		media.MediaFilename,
		media.MediaFilePath,
		media.CreatedAt,
	)

	return err
}

// Add Variation
func (s *PostgresStore) AddVariation(variation *Variation) error {
	query := `insert into media (
		id, 
		product_id, 
		variation_type, 
		variation_tag, 
		created_at) values ($1, $2, $3, $4, $5)`

	_, err := s.db.Query(
		query,
		variation.ID,
		variation.ProductID,
		variation.VariationType,
		variation.VariationTag,
		variation.CreatedAt,
	)

	return err
}

// Edit Product
func (s *PostgresStore) EditProduct(id int, product *Product) error {
	return nil
}

// Get Products
func (s *PostgresStore) GetProducts() ([]*Product, error) {
	query := `
		select
			p.id,
			p.thumbnail_name,
			p.thumbnail_path,
			p.status,
			p.category,
			p.tag,
			p.template,
			p.name,
			p.description,
			p.price,
			p.discount_type,
			p.tax_class,
			p.vat_amount,
			p.sku_number,
			p.barcode_number,
			p.on_shelf,
			p.on_warehouse,
			p.allow_backorder,
			p.in_physical,
			p.meta_title,
			p.meta_description,
			p.meta_keywords,
			to_char(p.created_at, 'YYYY-MM-DD"T"HH24:MI:SS.US"Z"') AS created_at,

			coalesce(jsonb_agg(distinct jsonb_build_object(
				'id', v.id,
				'variation_type', v.variation_type,
				'variation_tag', v.variation_tag,
				'created_at', to_char(v.created_at, 'YYYY-MM-DD"T"HH24:MI:SS.US"Z"')
			)) filter (where v.id is not null), '[]'::jsonb) as variations,

			coalesce(jsonb_agg(distinct jsonb_build_object(
				'id', m.id,
				'media_filename', m.media_filename,
				'media_file_path', m.media_file_path,
				'created_at', to_char(m.created_at, 'YYYY-MM-DD"T"HH24:MI:SS.US"Z"')
			)) filter (where m.id is not null), '[]'::jsonb) as media
		
		from product p
		left join variation v on p.id = v.product_id
		left join media m on p.id = m.product_id
		group by p.id
	`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := []*Product{}
	for rows.Next() {

		product, variationsJson, mediaJson, createdAtStr, err := scanIntoProduct(rows)
		if err != nil {
			return nil, err
		}

		// Parse created_at manually
		parsedTime, err := time.Parse("2006-01-02T15:04:05.999999Z", createdAtStr)
		if err != nil {
			return nil, err
		}
		product.CreatedAt = parsedTime.Format(time.RFC3339)

		// unmarshal variations
		if err := json.Unmarshal(variationsJson, &product.Variations); err != nil {
			return nil, err
		}

		// unmarshal media
		if err := json.Unmarshal(mediaJson, &product.Media); err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	return products, nil
}

// Get Medias
func (s *PostgresStore) GetMedias() ([]*Media, error) {
	query := "select * from media"
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	medias := []*Media{}
	for rows.Next() {
		media, createdAtStr, err := scanIntoMedia(rows)
		if err != nil {
			return nil, err
		}

		// parse createdAt
		parsedTime, err := time.Parse("2006-01-02T15:04:05.999999Z", createdAtStr)
		if err != nil {
			return nil, err
		}
		media.CreatedAt = parsedTime.Format(time.RFC3339)

		medias = append(medias, media)
	}

	return medias, nil
}

// Get Product by id
func (s *PostgresStore) GetProductByID(id int) (*Product, error) {
	return nil, nil
}

// Delete Product
func (s *PostgresStore) DeleteProduct(id int) error {
	return nil
}

func scanIntoProduct(rows *sql.Rows) (*Product, []byte, []byte, string, error) {
	product := new(Product)
	var variationsJson, mediaJson []byte
	var createdAtStr string

	err := rows.Scan(
		&product.ID,
		&product.ThumbnailName,
		&product.ThumbnailPath,
		&product.Status,
		&product.Category,
		&product.Tag,
		&product.Template,
		&product.Name,
		&product.Description,
		&product.Price,
		&product.DiscountType,
		&product.TaxClass,
		&product.VATAmount,
		&product.SKUNumber,
		&product.BarcodeNumber,
		&product.OnShelf,
		&product.OnWarehouse,
		&product.AllowBackOrder,
		&product.InPhysical,
		&product.MetaTitle,
		&product.MetaDescription,
		&product.MetaKeywords,
		&createdAtStr,
		&variationsJson,
		&mediaJson,
	)

	return product, variationsJson, mediaJson, createdAtStr, err
}

func scanIntoMedia(rows *sql.Rows) (*Media, string, error) {
	media := new(Media)
	var createdAtStr string

	err := rows.Scan(
		&media.ID,
		&media.ProductID,
		&media.MediaFilename,
		&media.MediaFilePath,
		&createdAtStr,
	)

	return media, createdAtStr, err
}
