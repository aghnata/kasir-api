package entity

import "errors"

// Domain errors
var (
	ErrProductNameRequired = errors.New("product name is required")
	ErrProductPriceInvalid = errors.New("product price must be greater than zero")
	ErrProductStockInvalid = errors.New("product stock cannot be negative")
	ErrProductNotFound     = errors.New("product not found")
)

// Product represents the domain entity for a product
type Product struct {
	ID    int
	Name  string
	Price int
	Stock int
}

// Validate validates the product entity
func (p *Product) Validate() error {
	if p.Name == "" {
		return ErrProductNameRequired
	}
	if p.Price <= 0 {
		return ErrProductPriceInvalid
	}
	if p.Stock < 0 {
		return ErrProductStockInvalid
	}
	return nil
}

// NewProduct creates a new product with validation
func NewProduct(name string, price, stock int) (*Product, error) {
	p := &Product{
		Name:  name,
		Price: price,
		Stock: stock,
	}
	if err := p.Validate(); err != nil {
		return nil, err
	}
	return p, nil
}
