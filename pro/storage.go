package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateProduct(*Product) error
	GetProducts() (*[]Product, error)
	GetProductByID(int) (*Product, error)
	UpdateProduct(int, *Product) error
	DeleteProduct(int) error
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
		id serial primary key,
		name varchar(100) not null,
		brand varchar(100),
		category varchar(50),
		price integer not null,
		quantity integer not null,
		created_at timestamp
	)`

	_, err := s.db.Exec(query)
	return err
}

func (s *PostgresStore) CreateProduct(product *Product) error {
	return nil
}

func (s *PostgresStore) GetProducts() (*[]Product, error) {
	return nil, nil
}

func (s *PostgresStore) GetProductByID(id int) (*Product, error) {
	return nil, nil
}

func (s *PostgresStore) DeleteProduct(id int) error {
	return nil
}

func (s *PostgresStore) UpdateProduct(id int, product *Product) error {
	return nil
}
