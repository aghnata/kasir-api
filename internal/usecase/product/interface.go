package product

// ProductUseCase defines the interface for product business logic
type ProductUseCase interface {
	// GetAll retrieves all products
	GetAll() ([]ProductOutput, error)

	// GetByID retrieves a product by its ID
	GetByID(id int) (*ProductOutput, error)

	// Create creates a new product
	Create(input CreateProductInput) (*ProductOutput, error)

	// Update updates an existing product
	Update(id int, input UpdateProductInput) (*ProductOutput, error)

	// Delete deletes a product by its ID
	Delete(id int) error
}
