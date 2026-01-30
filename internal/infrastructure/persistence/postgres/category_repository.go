package postgres

import (
	"database/sql"

	"kasir-api/internal/domain/entity"
)

// CategoryRepository implements repository.CategoryRepository interface
type CategoryRepository struct {
	db *sql.DB
}

// NewCategoryRepository creates a new CategoryRepository
func NewCategoryRepository(db *sql.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

// GetAll retrieves all categories from the database
func (r *CategoryRepository) GetAll() ([]entity.Category, error) {
	query := "SELECT id, name, description FROM categories"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	categories := make([]entity.Category, 0)
	for rows.Next() {
		var category entity.Category
		err := rows.Scan(&category.ID, &category.Name, &category.Description)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return categories, nil
}

// GetByID retrieves a category by its ID
func (r *CategoryRepository) GetByID(id int) (*entity.Category, error) {
	query := "SELECT id, name, description FROM categories WHERE id = $1"
	row := r.db.QueryRow(query, id)

	var category entity.Category
	err := row.Scan(&category.ID, &category.Name, &category.Description)
	if err == sql.ErrNoRows {
		return nil, entity.ErrCategoryNotFound
	}
	if err != nil {
		return nil, err
	}

	return &category, nil
}

// Create creates a new category in the database
func (r *CategoryRepository) Create(category *entity.Category) error {
	query := "INSERT INTO categories (name, description) VALUES ($1, $2) RETURNING id"
	err := r.db.QueryRow(query, category.Name, category.Description).Scan(&category.ID)
	if err != nil {
		return err
	}
	return nil
}

// Update updates an existing category in the database
func (r *CategoryRepository) Update(category *entity.Category) error {
	query := "UPDATE categories SET name = $1, description = $2 WHERE id = $3"
	result, err := r.db.Exec(query, category.Name, category.Description, category.ID)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return entity.ErrCategoryNotFound
	}

	return nil
}

// Delete deletes a category from the database
func (r *CategoryRepository) Delete(id int) error {
	query := "DELETE FROM categories WHERE id = $1"
	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return entity.ErrCategoryNotFound
	}

	return nil
}
