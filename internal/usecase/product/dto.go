package product

// CreateProductInput represents the input for creating a product
type CreateProductInput struct {
	Name  string
	Price int
	Stock int
}

// UpdateProductInput represents the input for updating a product
type UpdateProductInput struct {
	Name  string
	Price int
	Stock int
}

// ProductOutput represents the output of product operations
type ProductOutput struct {
	ID    int
	Name  string
	Price int
	Stock int
}
