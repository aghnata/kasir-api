package category

// CategoryUseCase defines the interface for category business logic
type CategoryUseCase interface {
	// GetAll retrieves all categories
	GetAll() ([]CategoryOutput, error)

	// GetByID retrieves a category by its ID
	GetByID(id int) (*CategoryOutput, error)

	// Create creates a new category
	Create(input CreateCategoryInput) (*CategoryOutput, error)

	// Update updates an existing category
	Update(id int, input UpdateCategoryInput) (*CategoryOutput, error)

	// Delete deletes a category by its ID
	Delete(id int) error
}
