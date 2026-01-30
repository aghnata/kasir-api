package entity

import "errors"

// Domain errors
var (
	ErrCategoryNameRequired = errors.New("category name is required")
	ErrCategoryNotFound     = errors.New("category not found")
)

// Category represents the domain entity for a category
type Category struct {
	ID          int
	Name        string
	Description string
}

// Validate validates the category entity
func (c *Category) Validate() error {
	if c.Name == "" {
		return ErrCategoryNameRequired
	}
	return nil
}

// NewCategory creates a new category with validation
func NewCategory(name, description string) (*Category, error) {
	c := &Category{
		Name:        name,
		Description: description,
	}
	if err := c.Validate(); err != nil {
		return nil, err
	}
	return c, nil
}
