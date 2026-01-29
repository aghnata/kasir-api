package repositories

import (
	"database/sql"
	"errors"
	"kasir-api/models"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (repo *ProductRepository) GetAll() ([]models.Product, error) {
	query := "SELECT id, name, price, stock FROM products"
	rows, err := repo.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := make([]models.Product, 0)
	for rows.Next() {
		var product models.Product
		err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.Stock)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

func (repo *ProductRepository) Create(product *models.Product) error {
	// Validate input
	if product == nil {
		return sql.ErrNoRows
	}

	// Additional validation could be added here
	// e.g., check if name is not empty, price is positive, etc.

	// Use parameterized query to prevent SQL injection
	query := "INSERT INTO products (name, price, stock) VALUES ($1, $2, $3) RETURNING id"
	err := repo.db.QueryRow(query, product.Name, product.Price, product.Stock).Scan(&product.ID)
	if err != nil {
		return err
	}

	return nil
}

func (repo *ProductRepository) GetByID(id int) (*models.Product, error) {
	query := "SELECT id, name, price, stock FROM products WHERE id = $1"
	row := repo.db.QueryRow(query, id)
	var product models.Product
	err := row.Scan(&product.ID, &product.Name, &product.Price, &product.Stock)

	if err == sql.ErrNoRows {
		return nil, errors.New("produk tidak ditemukan")
	}
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (repo *ProductRepository) Update(data *models.Product) error {
	// Implementation for updating a product
	query := "UPDATE products SET name = $1, price = $2, stock = $3 WHERE id = $4"
	result, err := repo.db.Exec(query, data.Name, data.Price, data.Stock, data.ID)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("produk tidak ditemukan")
	}

	return nil
}
