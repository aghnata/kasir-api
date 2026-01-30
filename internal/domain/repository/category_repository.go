package repository

import "kasir-api/internal/domain/entity"

// CategoryRepository defines the interface for category data access
// This interface is defined in the domain layer, allowing the domain
// to specify what it needs without depending on infrastructure details
type CategoryRepository interface {
	// GetAll retrieves all categories
	GetAll() ([]entity.Category, error)

	// GetByID retrieves a category by its ID
	GetByID(id int) (*entity.Category, error)

	// Create creates a new category
	Create(category *entity.Category) error

	// Update updates an existing category
	Update(category *entity.Category) error

	// Delete deletes a category by its ID
	Delete(id int) error
}
