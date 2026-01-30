package product

import (
	"kasir-api/internal/domain/entity"
	"kasir-api/internal/domain/repository"
)

// Service implements ProductUseCase interface
type Service struct {
	repo repository.ProductRepository
}

// NewService creates a new product service
func NewService(repo repository.ProductRepository) *Service {
	return &Service{repo: repo}
}

// GetAll retrieves all products
func (s *Service) GetAll() ([]ProductOutput, error) {
	products, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	outputs := make([]ProductOutput, len(products))
	for i, p := range products {
		outputs[i] = toProductOutput(&p)
	}
	return outputs, nil
}

// GetByID retrieves a product by its ID
func (s *Service) GetByID(id int) (*ProductOutput, error) {
	product, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	output := toProductOutput(product)
	return &output, nil
}

// Create creates a new product with validation
func (s *Service) Create(input CreateProductInput) (*ProductOutput, error) {
	// Use domain entity for validation
	product, err := entity.NewProduct(input.Name, input.Price, input.Stock)
	if err != nil {
		return nil, err
	}

	err = s.repo.Create(product)
	if err != nil {
		return nil, err
	}

	output := toProductOutput(product)
	return &output, nil
}

// Update updates an existing product
func (s *Service) Update(id int, input UpdateProductInput) (*ProductOutput, error) {
	// Check if product exists
	existing, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	// Update fields
	existing.Name = input.Name
	existing.Price = input.Price
	existing.Stock = input.Stock

	// Validate
	if err := existing.Validate(); err != nil {
		return nil, err
	}

	err = s.repo.Update(existing)
	if err != nil {
		return nil, err
	}

	output := toProductOutput(existing)
	return &output, nil
}

// Delete deletes a product by its ID
func (s *Service) Delete(id int) error {
	return s.repo.Delete(id)
}

// toProductOutput converts domain entity to output DTO
func toProductOutput(p *entity.Product) ProductOutput {
	return ProductOutput{
		ID:    p.ID,
		Name:  p.Name,
		Price: p.Price,
		Stock: p.Stock,
	}
}
