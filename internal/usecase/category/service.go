package category

import (
	"kasir-api/internal/domain/entity"
	"kasir-api/internal/domain/repository"
)

// Service implements CategoryUseCase interface
type Service struct {
	repo repository.CategoryRepository
}

// NewService creates a new category service
func NewService(repo repository.CategoryRepository) *Service {
	return &Service{repo: repo}
}

// GetAll retrieves all categories
func (s *Service) GetAll() ([]CategoryOutput, error) {
	categories, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	outputs := make([]CategoryOutput, len(categories))
	for i, c := range categories {
		outputs[i] = toCategoryOutput(&c)
	}
	return outputs, nil
}

// GetByID retrieves a category by its ID
func (s *Service) GetByID(id int) (*CategoryOutput, error) {
	category, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	output := toCategoryOutput(category)
	return &output, nil
}

// Create creates a new category with validation
func (s *Service) Create(input CreateCategoryInput) (*CategoryOutput, error) {
	// Use domain entity for validation
	category, err := entity.NewCategory(input.Name, input.Description)
	if err != nil {
		return nil, err
	}

	err = s.repo.Create(category)
	if err != nil {
		return nil, err
	}

	output := toCategoryOutput(category)
	return &output, nil
}

// Update updates an existing category
func (s *Service) Update(id int, input UpdateCategoryInput) (*CategoryOutput, error) {
	// Check if category exists
	existing, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	// Update fields
	existing.Name = input.Name
	existing.Description = input.Description

	// Validate
	if err := existing.Validate(); err != nil {
		return nil, err
	}

	err = s.repo.Update(existing)
	if err != nil {
		return nil, err
	}

	output := toCategoryOutput(existing)
	return &output, nil
}

// Delete deletes a category by its ID
func (s *Service) Delete(id int) error {
	return s.repo.Delete(id)
}

// toCategoryOutput converts domain entity to output DTO
func toCategoryOutput(c *entity.Category) CategoryOutput {
	return CategoryOutput{
		ID:          c.ID,
		Name:        c.Name,
		Description: c.Description,
	}
}
