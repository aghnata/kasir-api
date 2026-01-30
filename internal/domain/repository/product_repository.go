package repository

import "kasir-api/internal/domain/entity"

// ProductRepository defines the interface for product data access
// This interface is defined in the domain layer, allowing the domain
// to specify what it needs without depending on infrastructure details
type ProductRepository interface {
	// GetAll retrieves all products
	GetAll() ([]entity.Product, error)

	// GetByID retrieves a product by its ID
	GetByID(id int) (*entity.Product, error)

	// Create creates a new product
	Create(product *entity.Product) error

	// Update updates an existing product
	Update(product *entity.Product) error

	// Delete deletes a product by its ID
	Delete(id int) error
}
