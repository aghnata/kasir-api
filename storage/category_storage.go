package storage

import (
	"errors"
	"kasir-api/models"
)

// CategoryStorage handles category data operations
type CategoryStorage struct {
	category []models.Category
}

// NewCategoryStorage creates a new CategoryStorage instance with initial data
func NewCategoryStorage() *CategoryStorage {
	return &CategoryStorage{
		category: []models.Category{
			{ID: 1, Name: "Category A", Description: "Description A"},
			{ID: 2, Name: "Category B", Description: "Description B"},
			{ID: 3, Name: "Category C", Description: "Description C"},
		},
	}
}

// GetAll returns all categories
func (cs *CategoryStorage) GetAll() []models.Category {
	return cs.category
}

// GetByID returns a category by ID
func (cs *CategoryStorage) GetByID(id int) (*models.Category, error) {
	for _, c := range cs.category {
		if c.ID == id {
			return &c, nil
		}
	}
	return nil, errors.New("kategori tidak ditemukan")
}

// Create adds a new category
func (cs *CategoryStorage) Create(categoryBaru models.Category) models.Category {
	categoryBaru.ID = len(cs.category) + 1
	cs.category = append(cs.category, categoryBaru)
	return categoryBaru
}

// Update updates an existing category by ID
func (cs *CategoryStorage) Update(id int, categoryUpdate models.Category) (*models.Category, error) {
	for i, c := range cs.category {
		if c.ID == id {
			categoryUpdate.ID = id
			cs.category[i] = categoryUpdate
			return &categoryUpdate, nil
		}
	}
	return nil, errors.New("kategori tidak ditemukan")
}

// Delete removes a category by ID
func (cs *CategoryStorage) Delete(id int) error {
	for i, c := range cs.category {
		if c.ID == id {
			cs.category = append(cs.category[:i], cs.category[i+1:]...)
			return nil
		}
	}
	return errors.New("kategori tidak ditemukan")
}
