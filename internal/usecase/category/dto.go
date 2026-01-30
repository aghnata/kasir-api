package category

// CreateCategoryInput represents the input for creating a category
type CreateCategoryInput struct {
	Name        string
	Description string
}

// UpdateCategoryInput represents the input for updating a category
type UpdateCategoryInput struct {
	Name        string
	Description string
}

// CategoryOutput represents the output of category operations
type CategoryOutput struct {
	ID          int
	Name        string
	Description string
}
