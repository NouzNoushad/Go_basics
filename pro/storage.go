package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateProduct(*Product) error
	GetProducts() ([]*Product, error)
	GetProductByID(string) (*Product, error)
	UpdateProduct(string, *Product) error
	DeleteProduct(string) error
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	connStr := "host=localhost user=postgres password=noushad dbname=pro port=5432 sslmode=disable"
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

func (s *PostgresStore) Init() error {

	return s.createAccountTable()
}

func (s *PostgresStore) createAccountTable() error {
	query := `create table if not exists product(
		id varchar(100) primary key,
		name varchar(100) not null,
		brand varchar(100),
		category varchar(50),
		price numeric not null,
		quantity integer not null,
		created_at timestamp
	)`

	_, err := s.db.Exec(query)
	return err
}

func (s *PostgresStore) CreateProduct(product *Product) error {
	query := `insert into product (id, name, brand, category, price, quantity, created_at) values ($1, $2, $3, $4, $5, $6, $7)`

	_, err := s.db.Query(
		query,
		product.ID,
		product.Name,
		product.Brand,
		product.Category,
		product.Price,
		product.Quantity,
		product.CreatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *PostgresStore) GetProducts() ([]*Product, error) {
	query := "select * from product"
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}

	products := []*Product{}
	for rows.Next() {
		product, err := scanIntoAccount(rows)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

func scanIntoAccount(rows *sql.Rows) (*Product, error) {
	product := new(Product)
	err := rows.Scan(
		&product.ID,
		&product.Name,
		&product.Brand,
		&product.Category,
		&product.Price,
		&product.Quantity,
		&product.CreatedAt)

	return product, err
}

func (s *PostgresStore) GetProductByID(id string) (*Product, error) {
	query := "select * from product where id = $1"
	row, err := s.db.Query(query, id)
	if err != nil {
		return nil, err
	}

	for row.Next() {
		return scanIntoAccount(row)
	}

	return nil, fmt.Errorf("account with id [%s] not found", id)
}

func (s *PostgresStore) DeleteProduct(id string) error {
	query := "delete from product where id = $1"
	_, err := s.db.Query(query, id)

	return err
}

func (s *PostgresStore) UpdateProduct(id string, product *Product) error {
	query := "update product set name=$1, brand=$2, category=$3, price=$4, quantity=$5 where id=$6"
	_, err := s.db.Query(
		query,
		product.Name,
		product.Brand,
		product.Category,
		product.Price,
		product.Quantity,
		id)

	if err != nil {
		return err
	}

	return nil
}
